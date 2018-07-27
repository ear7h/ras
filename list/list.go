/*
NOTE: in most functions, the addresses passed around
are that of the usable storage, that is after the allocated
block header
 */
package list

import (
	"fmt"
	"errors"
	"os"
	"encoding/binary"
	"sync"
	"sync/atomic"
)

const NullPtr uint = 0
const Debug = true

const (
	ioRead  = iota
	ioWrite
	ioTrunc
)

const (
	manAlloc  = iota
	manFree
	manResize
)

type manOp struct {
	op  int
	arg uint // addr or size
}

type manRes struct {
	val uint // addr for
}

type Allocator struct {
	Filename string

	cap   uint // max length of file
	file  atomic.Value
	flock sync.RWMutex
	mlock sync.Mutex // lock for free and alloc
}

func (h *Allocator) f() (f *os.File, err error) {
	v := h.file.Load()

	if v == nil {
		h.flock.Lock()
		defer h.flock.Unlock()

		var err error
		f, err = h.openFile()
		if err != nil {
			return nil, err
		}
		h.file.Store(f)
	} else {
		f = v.(*os.File)
	}

	return f, nil
}

// Cap returns the maximum length of the file
func (h *Allocator) Cap() uint {
	if h.file.Load() == nil {
		h.f()
	}

	return h.cap
}

func (h *Allocator) openFile() (f *os.File, err error) {

	f, err = os.OpenFile(h.Filename, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return f, err
	}

	stat, err := f.Stat()
	if err != nil {
		return f, err
	}

	h.cap = uint(stat.Size())

	if h.cap == 0 {
		hdrByt := make([]byte, usableOffset)
		copy(hdrByt[magicOffset:magicOffset+magicLen],
			Magic[:])

		_, err = f.WriteAt(hdrByt, 0)
		if err != nil {
			return f, err
		}

		h.cap = usableOffset
	} else if h.cap < usableOffset {
		return f, fmt.Errorf("file corrupted, size not 0 or big enough for header (%d)", usableOffset)
	} else {
		hdr := make([]byte, magicLen)
		_, err = f.ReadAt(hdr, 0)
		if err != nil {
			return f, err
		}

		for i := 0; i < magicLen; i++ {
			if hdr[i] != Magic[i] {
				return f, fmt.Errorf("unrecognized magic numbers")
			}
		}
	}

	fmt.Println("cap at openFile:", h.cap)

	return f, nil
}

var trBreakIter = errors.New("break iter")
var trEndIter = errors.New("end iter")

func (h Allocator) traverse(fn func(addr uint, blk blockHeader) error) error {
	var hdr blockHeader
	var err error

	headPtr, err := h.getHeadPtr()
	if err != nil {
		return err
	}

	for addr := headPtr; addr != NullPtr; addr = hdr.next {
		hdr, err = h.getHdr(addr)
		if err != nil {
			return err
		}

		err = fn(addr, hdr)
		if err != nil {
			return err
		}
	}

	return trEndIter
}

func (h Allocator) setHeadPtr(addr uint) error {
	byt := make([]byte, headPtrLen)
	binary.LittleEndian.PutUint32(byt, uint32(addr))
	return h.Set(headPtrOffset, byt)
}

func (h Allocator) getHeadPtr() (uint, error) {
	byt := make([]byte, headPtrLen)
	err := h.Get(headPtrOffset, byt)
	if err != nil {
		return 0, err
	}

	return uint(binary.LittleEndian.Uint32(byt)), nil
}

// sets the header corresponding with the address
func (h Allocator) setHdr(addr uint, blk blockHeader) error {
	fmt.Println("setting ", addr, blk)
	return h.Set(addr-blockHeaderLen, blk.MarshalBinary())
}

// gets the header corresponding with the address
func (h Allocator) getHdr(addr uint) (blockHeader, error) {
	dst := make([]byte, blockHeaderLen)

	err := h.Get(addr-blockHeaderLen, dst)
	if err != nil {
		return blockHeader{}, err
	}

	ret := blockHeader{}
	ret.UnmarshalBinary(dst)

	return ret, nil
}

// addr is the address of a storage block
// size is the size of the new storage block
func (h Allocator) insertBlockAfter(addr, size uint) error {
	fmt.Println("inseting block after ", addr, " size ", size)
	if addr == usableOffset {
		next, err := h.getHeadPtr()
		if err != nil {
			return err
		}

		err = h.setHeadPtr(usableOffset + blockHeaderLen)
		if err != nil {
			return err
		}

		hdr2 := blockHeader{
			len:  size,
			next: next,
		}

		err = h.setHdr(usableOffset+blockHeaderLen, hdr2)
		if err != nil {
			return err
		}

		return nil
	}

	hdr1, err := h.getHdr(addr)
	if err != nil {
		return err
	}

	hdr2 := blockHeader{
		len:  size,
		next: hdr1.next,
	}

	hdr1.next = addr + hdr1.len + blockHeaderLen

	err = h.setHdr(addr, hdr1)
	if err != nil {
		return err
	}

	err = h.setHdr(hdr1.next, hdr2)
	return nil
}

// addr is the address of a storage block
func (h Allocator) removeBlockAfter(addr uint) error {
	if addr == usableOffset {
		hdr, err := h.getHdr(usableOffset + blockHeaderLen)
		if err != nil {
			return err
		}

		return h.setHeadPtr(hdr.next)
	}

	hdr, err := h.getHdr(addr)
	if err != nil {
		return err
	}

	hdr2, err := h.getHdr(hdr.next)
	if err != nil {
		return err
	}

	hdr.next = hdr2.next

	return h.setHdr(addr, hdr)
}

func (h *Allocator) tsafeManOp(op int, arg uint) (r uint, err error) {
	h.mlock.Lock()
	defer h.mlock.Unlock()
	var str string
	switch op {
	case manAlloc:
		str = "alloc"
	case manFree:
		str = "free"
	case manResize:
		str = "resize"
	}

	fmt.Println("man op ", str, "arg: ", arg)

	switch op {
	case manAlloc:
		r, err = h.alloc(arg)
	case manFree:
		err = h.free(arg)
	case manResize:
		err = h.resize(arg)
	}

	return
}

func (h Allocator) Alloc(size uint) (addr uint, err error) {
	addr, err = h.tsafeManOp(manAlloc, size)
	return
}

func (h Allocator) alloc(size uint) (addr uint, err error) {

	var pend uint = usableOffset
	var p uint = usableOffset

	err = h.traverse(func(caddr uint, blk blockHeader) error {

		// 2 headers need fit between the end of the last block
		// and the address of the current block
		// 1) the new block's header
		// 2) the current block's header
		if caddr-pend >= size+blockHeaderLen2 {
			addr = pend + blockHeaderLen
			err = h.insertBlockAfter(p, size)
			if err != nil {
				return err
			}

			return trBreakIter
		}

		p = caddr
		pend = caddr + blk.len

		return nil
	})

	switch err {
	case trBreakIter:
		return addr, nil

	case trEndIter:
		// this should only happen on the first allocation
		// where a.traverse() does no iterations and the
		// p and pend variaables are not mutated
		// and there is space to allocate
		if pend == usableOffset && h.cap >= p+size {
			goto L
		}

		// check for space at the end of the list
		if h.cap-pend < size+blockHeaderLen {
			return NullPtr, fmt.Errorf("no space e%d:c%d:s%d+%d", pend, h.cap, size, blockHeaderLen)
		}
	L:
		fmt.Println("pend", pend)
		addr = pend + blockHeaderLen
		err = h.insertBlockAfter(p, size)
		return addr, err

	default:
		return NullPtr, err
	}

}

func (h Allocator) Free(addr uint) (err error) {
	_, err = h.tsafeManOp(manFree, addr)
	return
}

func (h Allocator) free(addr uint) (err error) {

	var prev uint = usableOffset

	err = h.traverse(func(caddr uint, blk blockHeader) error {
		end := caddr + blk.len - 1

		if caddr <= addr && addr <= end {
			if caddr == addr {
				err = h.removeBlockAfter(prev)
				if err != nil {
					return err
				}
			} else {
				blk.len = addr - caddr
				err = h.setHdr(caddr, blk)
				if err != nil {
					return err
				}
			}

			return trBreakIter
		}

		prev = caddr
		return nil
	})

	switch err {
	case trBreakIter:
		return nil
	case trEndIter:
		return fmt.Errorf("free invalid address")
	default:
		return err
	}
}

func (h *Allocator) Resize(cap uint) (err error) {
	_, err = h.tsafeManOp(manResize, cap)
	return
}

func (h *Allocator) resize(cap uint) error {
	endsBefore, seg, err := h.endsBefore(cap)
	if err != nil {
		return err
	}

	if !endsBefore {
		return fmt.Errorf("shrink over allocared area %d", seg)
	}

	return h.tsafeIoOp(ioTrunc, nil, cap)
}

func (h Allocator) endsBefore(cap uint) (bool, uint, error) {
	if cap > h.cap {
		return true, h.cap, nil
	}

	if cap < usableOffset {
		return false, usableOffset, nil
	}

	var seg uint = usableOffset
	err := h.traverse(func(addr uint, blk blockHeader) error {
		if addr+blk.len > cap + 1{
			seg = addr+blk.len - 1
			return trBreakIter
		}

		return nil
	})

	if err == trBreakIter {
		return false, seg, nil
	} else if err == trEndIter {
		return true, seg, nil
	} else {
		return false, seg, err
	}

}

var i int32 = 0

func (h *Allocator) tsafeIoOp(op int, data []byte, arg uint) error {

	if Debug {
		ii := atomic.AddInt32(&i, 1)
		fmt.Println(ii, ") io op ", op)
		defer fmt.Println(ii, ") io op end ", op)
	}

	f, err := h.f()
	if err != nil {
		return err
	}

	switch op {
	case ioRead:
		h.flock.RLock()
		defer h.flock.RUnlock()

		_, err = f.ReadAt(data, int64(arg))

		return err

	case ioWrite:
		h.flock.Lock()
		defer h.flock.Unlock()

		_, err = f.WriteAt(data, int64(arg))
		if err != nil {
			return err
		}

		err = f.Sync()
		if err != nil {
			return err
		}

		return nil

	case ioTrunc:
		h.flock.Lock()
		defer h.flock.Unlock()

		err = f.Truncate(int64(arg))
		if err != nil {
			return err
		}

		err = f.Sync()
		if err != nil {
			return err
		}

		h.cap = arg

		return nil
	default:
		panic("iop not recognized")
	}
}

func (h *Allocator) Get(addr uint, dst []byte) (err error) {
	// todo: remove
	if Debug {
		defer func() {
			if err != nil {
				fmt.Println(addr, dst)
				panic(err)
			}
		}()
	}

	if addr+uint(len(dst)) > h.Cap() {
		return fmt.Errorf("get invalid address %d:%d:%d", addr, addr+uint(len(dst)), h.cap)
	}

	return h.tsafeIoOp(ioRead, dst, addr)
}

func (h *Allocator) Set(addr uint, src []byte) (err error) {
	// todo: remove
	if Debug {
		defer func() {
			if err != nil {
				fmt.Println(addr, src)
				panic(err)
			}
		}()
	}

	if addr+uint(len(src)) > h.Cap() {
		return fmt.Errorf("get invalid address %d:%d:%d", addr, addr+uint(len(src)), h.cap)
	}

	return h.tsafeIoOp(ioWrite, src, addr)
}

func (h Allocator) Bsize(addr uint) (size uint, err error) {
	err = h.traverse(func(caddr uint, blk blockHeader) error {
		if caddr == addr {
			size = blk.len
			return trBreakIter
		}
		return nil
	})

	if err == trBreakIter || err == trEndIter {
		err = nil
	}

	return size, err
}
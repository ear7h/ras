/*
NOTE: in most functions, the addresses passed around
are that of the usable storage, that is after the allocated
block header
 */
package ras

import (
	"fmt"
	"errors"
	"os"
	"encoding/binary"
)

type Heap struct {
	Filename string

	cap uint
	f   *os.File
}

func (h Heap) Cap() uint { return h.cap }

// block header
const blockHeaderLen = 8

// 2 block headers, used when allocating,
const blockHeaderLen2 = 2 * blockHeaderLen

// in the heap the block is stored as two little endian uint32s
// the first is the length of the block excluding the header
// the second is a pointer to the next block
type blockHeader struct {
	len, next uint
}

func (hdr *blockHeader) UnmarshalBinary(data []byte) error {
	if len(data) != blockHeaderLen {
		return fmt.Errorf("block header must be taken form %d bytes", blockHeaderLen)
	}

	hdr.len = uint(binary.LittleEndian.Uint32(data[:4]))
	hdr.next = uint(binary.LittleEndian.Uint32(data[4:]))

	return nil
}

func (hdr blockHeader) MarshalBinary() []byte {
	ret := make([]byte, blockHeaderLen)

	binary.LittleEndian.PutUint32(ret[:4], uint32(hdr.len))
	binary.LittleEndian.PutUint32(ret[4:], uint32(hdr.next))

	return ret
}

const heapHeaderLen = 2
const magicLen = 2
const heapFirstOffset = heapHeaderLen + blockHeaderLen

var Magic = [...]byte{0xe7, 0x89}

type heapHeader struct {
	magic [2]byte
}

func (h *Heap) Load() error {
	var err error
	h.f, err = os.OpenFile(h.Filename, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return err
	}

	stat, err := h.f.Stat()
	if err != nil {
		return err
	}

	h.cap = uint(stat.Size())
	if err != nil {
		return err
	}

	if h.cap == 0 {
		h.f.WriteAt(Magic[:], 0)
	} else {
		hdr := make([]byte, magicLen)
		_, err = h.f.ReadAt(hdr, 0)
		if err != nil {
			return err
		}

		for i := 0; i < magicLen; i++ {
			if hdr[i] != Magic[i] {
				return fmt.Errorf("unrecognized magic numbers")
			}
		}
	}

	return nil
}

var trBreakIter = errors.New("break iter")
var trEndIter = errors.New("end iter")

func (h *Heap) traverse(fn func(addr uint, blk blockHeader) error) error {
	var hdr blockHeader
	var err error

	var addr uint = heapFirstOffset
	hdr, err = h.getHdr(addr)
	if err != nil {
		return err
	}

	for hdr.len != 0 {
		err = fn(addr, hdr)
		if err != nil {
			return err
		}

		addr = hdr.next
		hdr, err = h.getHdr(addr)
		if err != nil {
			return err
		}
	}

	return trEndIter
}

// sets the header corresponding with the
func (h *Heap) setHdr(addr uint, blk blockHeader) error {
	return h.Set(addr-blockHeaderLen, blk.MarshalBinary())
}

func (h *Heap) getHdr(addr uint) (blockHeader, error) {
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
func (h *Heap) insertBlockAfter(addr, size uint) error {
	hdr1, err := h.getHdr(addr)
	if err != nil {
		return err
	}

	hdr2 := blockHeader{
		len:  size,
		next: hdr1.next,
	}

	hdr1.next = addr + hdr1.len

	err = h.setHdr(addr, hdr1)
	if err != nil {
		return err
	}

	err = h.setHdr(hdr1.next, hdr2)
	return nil
}

// addr is the address of a storage block
func (h *Heap) removeBlockAfter(addr uint) error {
	hdr, err := h.getHdr(addr)
	if err != nil {
		return err
	}

	hdr2, err := h.getHdr(hdr.next)
	if err != nil {
		return err
	}

	hdr.next = hdr2.next

	err = h.setHdr(addr, hdr)
	return err
}

func (h *Heap) Alloc(size uint) (addr uint, err error) {

	var pend uint = heapFirstOffset
	var p uint = heapFirstOffset

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
		pend = caddr + blk.len - 1

		return nil
	})

	if err == trEndIter {
		// this should only happen on the first allocation
		// where h.trverse() does no iterations and the
		// p and pend variaables are not mutated
		// and there is space to allocate
		if p == pend && h.cap-pend >= size {
			goto L
		}

		if h.cap-pend < size+blockHeaderLen {
			return addr, fmt.Errorf("no space e%d:c%d:s%d+%d", pend, h.cap, size, blockHeaderLen)
		}
	L:
		addr = pend + blockHeaderLen
		err = h.insertBlockAfter(p, size)

		return addr, err

	} else if err != nil {
		return addr, err
	}

	return addr, nil

}

func (h *Heap) Free(addr uint) (err error) {

	var prev uint = heapFirstOffset

	err = h.traverse(func(caddr uint, blk blockHeader) error {
		end := caddr - blk.len - 1

		if caddr <= addr && end >= addr {
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

	if err == trEndIter {
		return fmt.Errorf("free invalid address")
	} else if err != nil {
		return err
	}

	return nil
}

func (h *Heap) shrink(cap uint) error {
	err := h.traverse(func(addr uint, blk blockHeader) error {
		if addr+blk.len > cap {
			return trBreakIter
		}

		return nil
	})

	if err == trBreakIter {
		return fmt.Errorf("shrink over allocated areas")
	}

	return h.f.Truncate(int64(cap))
}

func (h *Heap) Resize(cap uint) error {
	if h.f == nil {
		err := h.Load()
		if err != nil {
			return err
		}
	}

	if cap < h.cap {
		err := h.shrink(cap)
		if err != nil {
			return err
		}

		h.cap = cap
		return nil
	}

	_, err := h.f.Seek(int64(cap-1), 0)
	if err != nil {
		return err
	}

	_, err = h.f.Write([]byte{0})
	if err != nil {
		return err
	}

	err = h.f.Sync()
	if err != nil {
		return err
	}

	h.cap = cap
	return nil
}

func (h *Heap) Set(addr uint, val []byte) (err error) {
	if h.f == nil {
		err = h.Load()
		if err != nil {
			return err
		}
	}

	if addr+uint(len(val)) > h.cap {
		return fmt.Errorf("set invalid address %d:%d", addr, addr+uint(len(val)))
	}

	_, err = h.f.WriteAt(val, int64(addr))
	if err == nil {
		err = h.f.Sync()
	}
	return
}

func (h *Heap) Get(addr uint, dst []byte) (err error) {
	if h.f == nil {
		err = h.Load()
		if err != nil {
			return err
		}
	}

	if addr+uint(len(dst)) > h.cap {
		return fmt.Errorf("get invalid address %d:%d", addr, addr+uint(len(dst)))
	}

	_, err = h.f.ReadAt(dst, int64(addr))
	return
}

type Uint8 struct {
	addr uint
	h    *Heap
}

func (v Uint8) Get() (uint8, error) {
	dst := make([]byte, 1)
	err := v.h.Get(v.addr, dst)
	return uint8(dst[0]), err
}

func (v Uint8) Set(val uint8) (error) {
	return v.h.Set(v.addr, []byte{byte(val)})
}

type Runtime struct {
	Filename string

	heap *Heap
}

func (r Runtime) NewUint8() (ret Uint8, err error) {
	addr, err := r.heap.Alloc(1)
	if err != nil {
		return ret, err
	}

	ret.addr = addr
	ret.h = r.heap

	return ret, err
}

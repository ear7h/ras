package ras

import "encoding/binary"

const sliceHeaderLen = 8

type Uint8Slice struct {
	addr uint
	a Allocator
}

func (slc Uint8Slice) Cap() (uint, error){
	return slc.a.Bsize(slc.addr)
}

func (slc Uint8Slice) Len() (uint, error) {
	data := make([]byte, 8)
	err := slc.a.Get(slc.addr, data)
	if err != nil {
		return 0, err
	}

	return uint(binary.
		LittleEndian.
		Uint64(data)), nil
}

func (slc Uint8Slice) GetI(i uint) (uint8, error) {
	l, err := slc.Len()
	if err != nil {
		return 0, err
	}
	if i >= l {
		return 0, ErrSliceOutOfBounds
	}

	addr := slc.addr + sliceHeaderLen + i

	return Uint8{addr:addr, a:slc.a}.Get()
}

func (slc Uint8Slice) SetI(i uint, v uint8) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	addr := slc.addr + sliceHeaderLen + i

	return Uint8{addr:addr, a:slc.a}.Set(v)
}
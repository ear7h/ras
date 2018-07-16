package ras

import "encoding/binary"

type Uint8Array struct {
	addr uint
	a Allocator
}

func (arr Uint8Array) Set(src []uint8, o uint) error {
	byt := make([]byte, len(src))
	for i, v := range src {
		byt[i] = byte(v)
	}

	return arr.a.Set(arr.addr + o, byt)
}

func (arr Uint8Array) Get(dst []uint8, o uint) (error) {
	byt := make([]byte, len(dst))

	err := arr.a.Get(arr.addr, byt)
	if err != nil {
		return err
	}

	for i, v := range byt {
		dst[i] = uint8(v)
	}

	return nil
}

type Uint16Array struct {
	addr uint
	a Allocator
}

func (arr Uint16Array) Set(src []uint16, o uint) error {
	byt := make([]byte, len(src) * 2)
	for i, v := range src {
		binary.
			LittleEndian.
			PutUint16(byt[i*2:(i+1)*2], v)
	}

	return arr.a.Set(arr.addr + o*2, byt)
}

func (arr Uint16Array) Get(dst []uint16, o uint) (error) {
	byt := make([]byte, len(dst) * 2)

	err := arr.a.Get(arr.addr + o*2, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = binary.
			LittleEndian.
			Uint16(byt[i*2:(i+1)*2])
	}

	return nil
}

type Int8Array struct {
	addr uint
	a Allocator
}

func (arr Int8Array) Set(src []int8, o uint) error {
	byt := make([]byte, len(src))
	for i, v := range src {
		byt[i] = byte(v)
	}

	return arr.a.Set(arr.addr + o, byt)
}

func (arr Int8Array) Get(dst []int8, o uint) (error) {
	byt := make([]byte, len(dst))

	err := arr.a.Get(arr.addr, byt)
	if err != nil {
		return err
	}

	for i, v := range byt {
		dst[i] = int8(v)
	}

	return nil
}

type Int16Array struct {
	addr uint
	a Allocator
}

func (arr Int16Array) Set(src []int16, o uint) error {
	byt := make([]byte, len(src) * 2)
	for i, v := range src {
		binary.
			LittleEndian.
			PutUint16(byt[i*2:(i+1)*2], uint16(v))
	}

	return arr.a.Set(arr.addr + o*2, byt)
}

func (arr Int16Array) Get(dst []int16, o uint) (error) {
	byt := make([]byte, len(dst) * 2)

	err := arr.a.Get(arr.addr + o*2, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = int16(binary.
			LittleEndian.
			Uint16(byt[i*2:(i+1)*2]))
	}

	return nil
}
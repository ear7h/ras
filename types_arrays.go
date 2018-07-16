//this file is generated
package ras

import (
	"encoding/binary"
)

type UintArray struct {
	addr uint
	a Allocator
}

func (arr UintArray) Get(dst []uint, o uint) error {
	byt := make([]byte, len(dst) * 8)

	err := arr.a.Get(arr.addr + o*8, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = uint(binary.
			LittleEndian.
			Uint64(byt[i*8:(i+1)*8]))
	}

	return nil
}

func (arr UintArray) Set(src []uint, o uint) error {
	byt := make([]byte, len(src) * 8)
	for i, v := range src {
		binary.
			LittleEndian.
			PutUint64(byt[i*8:(i+1)*8], uint64(v))
	}

	return arr.a.Set(arr.addr + o*8, byt)
}

type Uint8Array struct {
	addr uint
	a Allocator
}

func (arr Uint8Array) Get(dst []uint8, o uint) error {
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

func (arr Uint8Array) Set(src []uint8, o uint) error {
	byt := make([]byte, len(src))
	for i, v := range src {
		byt[i] = byte(v)
	}

	return arr.a.Set(arr.addr + o, byt)
}

type Uint16Array struct {
	addr uint
	a Allocator
}

func (arr Uint16Array) Get(dst []uint16, o uint) error {
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

func (arr Uint16Array) Set(src []uint16, o uint) error {
	byt := make([]byte, len(src) * 2)
	for i, v := range src {
		binary.
			LittleEndian.
			PutUint16(byt[i*2:(i+1)*2], v)
	}

	return arr.a.Set(arr.addr + o*2, byt)

}

type Uint32Array struct {
	addr uint
	a Allocator
}

func (arr Uint32Array) Get(dst []uint32, o uint) error {
	byt := make([]byte, len(dst) * 4)

	err := arr.a.Get(arr.addr + o*4, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = binary.
			LittleEndian.
			Uint32(byt[i*4:(i+1)*4])
	}

	return nil
}

func (arr Uint32Array) Set(src []uint32, o uint) error {
	byt := make([]byte, len(src) * 4)
	for i, v := range src {
		binary.
			LittleEndian.
			PutUint32(byt[i*4:(i+1)*4], v)
	}

	return arr.a.Set(arr.addr + o*4, byt)

}

type Uint64Array struct {
	addr uint
	a Allocator
}

func (arr Uint64Array) Get(dst []uint64, o uint) error {
	byt := make([]byte, len(dst) * 8)

	err := arr.a.Get(arr.addr + o*8, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = binary.
			LittleEndian.
			Uint64(byt[i*8:(i+1)*8])
	}

	return nil
}

func (arr Uint64Array) Set(src []uint64, o uint) error {
	byt := make([]byte, len(src) * 8)
	for i, v := range src {
		binary.
			LittleEndian.
			PutUint64(byt[i*8:(i+1)*8], v)
	}

	return arr.a.Set(arr.addr + o*8, byt)

}

type IntArray struct {
	addr uint
	a Allocator
}

func (arr IntArray) Get(dst []int, o uint) error {
	byt := make([]byte, len(dst) * 8)

	err := arr.a.Get(arr.addr + o*8, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = int(binary.
			LittleEndian.
			Uint64(byt[i*8:(i+1)*8]))
	}

	return nil
}

func (arr IntArray) Set(src []int, o uint) error {
	byt := make([]byte, len(src) * 8)
	for i, v := range src {
		binary.
			LittleEndian.
			PutUint64(byt[i*8:(i+1)*8], uint64(v))
	}

	return arr.a.Set(arr.addr + o*8, byt)
}

type Int8Array struct {
	addr uint
	a Allocator
}

func (arr Int8Array) Get(dst []int8, o uint) error {
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

func (arr Int8Array) Set(src []int8, o uint) error {
	byt := make([]byte, len(src))
	for i, v := range src {
		byt[i] = byte(v)
	}

	return arr.a.Set(arr.addr + o, byt)
}

type Int16Array struct {
	addr uint
	a Allocator
}

func (arr Int16Array) Get(dst []int16, o uint) error {
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

func (arr Int16Array) Set(src []int16, o uint) error {
	byt := make([]byte, len(src) * 2)
	for i, v := range src {
		binary.
			LittleEndian.
			PutUint16(byt[i*2:(i+1)*2], uint16(v))
	}

	return arr.a.Set(arr.addr + o*2, byt)
}

type Int32Array struct {
	addr uint
	a Allocator
}

func (arr Int32Array) Get(dst []int32, o uint) error {
	byt := make([]byte, len(dst) * 4)

	err := arr.a.Get(arr.addr + o*4, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = int32(binary.
			LittleEndian.
			Uint32(byt[i*4:(i+1)*4]))
	}

	return nil
}

func (arr Int32Array) Set(src []int32, o uint) error {
	byt := make([]byte, len(src) * 4)
	for i, v := range src {
		binary.
			LittleEndian.
			PutUint32(byt[i*4:(i+1)*4], uint32(v))
	}

	return arr.a.Set(arr.addr + o*4, byt)
}

type Int64Array struct {
	addr uint
	a Allocator
}

func (arr Int64Array) Get(dst []int64, o uint) error {
	byt := make([]byte, len(dst) * 8)

	err := arr.a.Get(arr.addr + o*8, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = int64(binary.
			LittleEndian.
			Uint64(byt[i*8:(i+1)*8]))
	}

	return nil
}

func (arr Int64Array) Set(src []int64, o uint) error {
	byt := make([]byte, len(src) * 8)
	for i, v := range src {
		binary.
			LittleEndian.
			PutUint64(byt[i*8:(i+1)*8], uint64(v))
	}

	return arr.a.Set(arr.addr + o*8, byt)
}


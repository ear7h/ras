// this file is generated
package ras

import (
	"encoding/binary"
	"math"
)

type Uint struct {
	addr uint
	a Allocator
}

func (v Uint) Get() (uint, error) {
	dst := make([]byte, 8)
	err := v.a.Get(v.addr, dst)
	return uint(binary.LittleEndian.Uint64(dst)), err
}

func (v Uint) Set(val uint) error {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, uint64(val))
	return v.a.Set(v.addr, data)
}

type Uint8 struct {
	addr uint
	a Allocator
}

func (v Uint8) Get() (uint8, error) {
	dst := make([]byte, 1)
	err := v.a.Get(v.addr, dst)
	return uint8(dst[0]), err
}

func (v Uint8) Set(val uint8) error {
	return v.a.Set(v.addr, []byte{byte(val)})
}

type Uint16 struct {
	addr uint
	a Allocator
}

func (v Uint16) Get() (uint16, error) {
	dst := make([]byte, 2)
	err := v.a.Get(v.addr, dst)
	return binary.LittleEndian.Uint16(dst), err
}

func (v Uint16) Set(val uint16) error {
	data := make([]byte, 2)
	binary.LittleEndian.PutUint16(data, val)
	return v.a.Set(v.addr, data)
}

type Uint32 struct {
	addr uint
	a Allocator
}

func (v Uint32) Get() (uint32, error) {
	dst := make([]byte, 4)
	err := v.a.Get(v.addr, dst)
	return binary.LittleEndian.Uint32(dst), err
}

func (v Uint32) Set(val uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, val)
	return v.a.Set(v.addr, data)
}

type Uint64 struct {
	addr uint
	a Allocator
}

func (v Uint64) Get() (uint64, error) {
	dst := make([]byte, 8)
	err := v.a.Get(v.addr, dst)
	return binary.LittleEndian.Uint64(dst), err
}

func (v Uint64) Set(val uint64) error {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, val)
	return v.a.Set(v.addr, data)
}

type Int struct {
	addr uint
	a Allocator
}

func (v Int) Get() (int, error) {
	dst := make([]byte, 8)
	err := v.a.Get(v.addr, dst)
	return int(binary.LittleEndian.Uint64(dst)), err
}

func (v Int) Set(val int) error {
	data := make([]byte, 8)
		binary.LittleEndian.PutUint64(data, uint64(val))
		return v.a.Set(v.addr, data)
}

type Int8 struct {
	addr uint
	a Allocator
}

func (v Int8) Get() (int8, error) {
	dst := make([]byte, 1)
	err := v.a.Get(v.addr, dst)
	return int8(dst[0]), err
}

func (v Int8) Set(val int8) error {
	return v.a.Set(v.addr, []byte{byte(val)})
}

type Int16 struct {
	addr uint
	a Allocator
}

func (v Int16) Get() (int16, error) {
	dst := make([]byte, 2)
	err := v.a.Get(v.addr, dst)
	return int16(binary.LittleEndian.Uint16(dst)), err
}

func (v Int16) Set(val int16) error {
	data := make([]byte, 2)
	binary.LittleEndian.PutUint16(data, uint16(val))
	return v.a.Set(v.addr, data)
}

type Int32 struct {
	addr uint
	a Allocator
}

func (v Int32) Get() (int32, error) {
	dst := make([]byte, 4)
	err := v.a.Get(v.addr, dst)
	return int32(binary.LittleEndian.Uint32(dst)), err
}

func (v Int32) Set(val int32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, uint32(val))
	return v.a.Set(v.addr, data)
}

type Int64 struct {
	addr uint
	a Allocator
}

func (v Int64) Get() (int64, error) {
	dst := make([]byte, 8)
	err := v.a.Get(v.addr, dst)
	return int64(binary.LittleEndian.Uint64(dst)), err
}

func (v Int64) Set(val int64) error {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, uint64(val))
	return v.a.Set(v.addr, data)
}

type Float32 struct {
	addr uint
	a Allocator
}

func (v Float32) Get() (float32, error) {
	dst := make([]byte, 4)
	err := v.a.Get(v.addr, dst)
	ui := binary.LittleEndian.Uint32(dst)
	return math.Float32frombits(ui), err
}

func (v Float32) Set(val float32) error {
	data := make([]byte, 4)
	ui := math.Float32bits(val)
	binary.LittleEndian.PutUint32(data, ui)
	return v.a.Set(v.addr, data)
}

type Float64 struct {
	addr uint
	a Allocator
}

func (v Float64) Get() (float64, error) {
	dst := make([]byte, 8)
	err := v.a.Get(v.addr, dst)
	ui := binary.LittleEndian.Uint64(dst)
	return math.Float64frombits(ui), err
}

func (v Float64) Set(val float64) error {
	data := make([]byte, 8)
	ui := math.Float64bits(val)
	binary.LittleEndian.PutUint64(data, ui)
	return v.a.Set(v.addr, data)
}


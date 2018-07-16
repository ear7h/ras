package types_gen

import (
	"fmt"
	"reflect"
)

var BitSizeMap = map[reflect.Kind]uint{
	reflect.Uint: 0,
	reflect.Uint8: 8,
	reflect.Uint16: 16,
	reflect.Uint32: 32,
	reflect.Uint64: 64,
	reflect.Int: 0,
	reflect.Int8: 8,
	reflect.Int16: 16,
	reflect.Int32: 32,
	reflect.Int64: 64,
	reflect.Float32: 32,
	reflect.Float64: 64,
	reflect.Bool: 8,
}

func UintGet(bitSize uint) string {
	if bitSize == 0 {
		return `dst := make([]byte, 8)
	err := v.a.Get(v.addr, dst)
	return uint(binary.LittleEndian.Uint64(dst)), err`
	}

	if bitSize == 8 {
		return `dst := make([]byte, 1)
	err := v.a.Get(v.addr, dst)
	return uint8(dst[0]), err`
	}

	fmtStr := `dst := make([]byte, %d)
	err := v.a.Get(v.addr, dst)
	return binary.LittleEndian.Uint%d(dst), err`

	return fmt.Sprintf(fmtStr, bitSize / 8, bitSize)
}

func UintSet(bitSize uint) string {
	if bitSize == 0 {
		return `data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, uint64(val))
	return v.a.Set(v.addr, data)`
	}

	if bitSize == 8 {
		return `return v.a.Set(v.addr, []byte{byte(val)})`
	}

	fmtStr := `data := make([]byte, %d)
	binary.LittleEndian.PutUint%d(data, val)
	return v.a.Set(v.addr, data)`

	return fmt.Sprintf(fmtStr, bitSize/8, bitSize)
}

func IntGet(bitSize uint) string {
	if bitSize == 0 {
		return `dst := make([]byte, 8)
	err := v.a.Get(v.addr, dst)
	return int(binary.LittleEndian.Uint64(dst)), err`
	}

	if bitSize == 8 {
		return `dst := make([]byte, 1)
	err := v.a.Get(v.addr, dst)
	return int8(dst[0]), err`
	}

	fmtStr := `dst := make([]byte, %d)
	err := v.a.Get(v.addr, dst)
	return int%d(binary.LittleEndian.Uint%d(dst)), err`

	return fmt.Sprintf(fmtStr, bitSize / 8, bitSize, bitSize)
}

func IntSet(bitSize uint) string {
	if bitSize == 0 {
		return `data := make([]byte, 8)
		binary.LittleEndian.PutUint64(data, uint64(val))
		return v.a.Set(v.addr, data)`
	}

	if bitSize == 8 {
		return `return v.a.Set(v.addr, []byte{byte(val)})`
	}

	fmtStr := `data := make([]byte, %d)
	binary.LittleEndian.PutUint%d(data, uint%d(val))
	return v.a.Set(v.addr, data)`

	return fmt.Sprintf(fmtStr, bitSize/8, bitSize, bitSize)
}
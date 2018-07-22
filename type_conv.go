//+build none

package ras

import (
	"encoding/binary"
	"math"
)

func BytesToBool(byt []byte) bool {
	return byt[0] != 0
}

func BoolToBytes(b bool, byt []byte) {
	if b {
		byt[0] = 1
	} else {
		byt[0] = 0
	}
}

func BytesToUint8(byt []byte) uint8 {
	return uint8(byt[0])
}

func Uint8ToBytes(i uint8, byt []byte) {
	byt[0] = byte(i)
	return
}

func BytesToUint16(byt []byte) uint16 {
	return binary.LittleEndian.Uint16(byt)
}

func Uint16ToBytes(i uint16, byt []byte) {
	binary.LittleEndian.PutUint16(byt, i)
	return
}
func BytesToUint32(byt []byte) uint32 {
	return binary.LittleEndian.Uint32(byt)
}

func Uint32ToBytes(i uint32) []byte {
	byt := make([]byte, Uint32Size)
	binary.LittleEndian.PutUint32(byt, i)
	return byt
}
func BytesToUint64(byt []byte) uint64 {
	return binary.LittleEndian.Uint64(byt)
}

func Uint64ToBytes(i uint64) []byte {
	byt := make([]byte, Uint64Size)
	binary.LittleEndian.PutUint64(byt, i)
	return byt
}

// ints

func BytesToInt8(byt []byte) int8 {
	return int8(byt[0])
}

func Int8ToBytes(i int8) []byte {
	return []byte{byte(i)}
}

func BytesToInt16(byt []byte) int16 {
	return int16(BytesToUint16(byt))
}

func Int16ToBytes(i int16) []byte {
	byt := make([]byte, Int16Size)
	binary.LittleEndian.PutUint16(byt, uint16(i))
	return byt
}

func BytesToInt32(byt []byte) int32 {
	return int32(binary.LittleEndian.Uint32(byt))
}

func Int32ToBytes(i int32) []byte {
	byt := make([]byte, Int32Size)
	binary.LittleEndian.PutUint32(byt, uint32(i))
	return byt
}

func BytesToInt64(byt []byte) int64 {
	return int64(binary.LittleEndian.Uint64(byt))
}

func Int64ToBytes(i int64) []byte {
	byt := make([]byte, Int64Size)
	binary.LittleEndian.PutUint64(byt, uint64(i))
	return byt
}

// floats

func BytesToFloat32(byt []byte) float32 {
	i := binary.LittleEndian.Uint32(byt)
	return math.Float32frombits(i)
}

func Float32ToBytes(f float32) []byte {
	byt := make([]byte, Float32Size)
	i := math.Float32bits(f)
	binary.LittleEndian.PutUint32(byt, i)
	return byt
}

func BytesToFloat64(byt []byte) float64 {
	i := binary.LittleEndian.Uint64(byt)
	return math.Float64frombits(i)
}

func Float64ToBytes(f float64) []byte {
	byt := make([]byte, Float64Size)
	i := math.Float64bits(f)
	binary.LittleEndian.PutUint64(byt, i)
	return byt
}

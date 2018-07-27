
package ras

import (
	"encoding/binary"
	"math"
)


//
// uints
//
func BytesToUint8(byt []byte) uint8 {
	return uint8(byt[0])
}

func Uint8ToBytes(i uint8, byt []byte) {
	byt[0] = byte(i)
}

func BytesToUint(byt []byte) uint {
	return uint(BytesToUint64(byt))
}

func UintToBytes(i uint, byt []byte) {
	Uint64ToBytes(uint64(i), byt)
}

//go:proto T=Uint16,Uint32,Uint64 B=T.lower S=T.sizebits
func BytesToUint16(byt []byte) uint16 {
	return binary.LittleEndian.Uint16(byt)
}

func Uint16ToBytes(i uint16, byt []byte) {
	binary.LittleEndian.PutUint16(byt, i)
}
func BytesToUint32(byt []byte) uint32 {
	return binary.LittleEndian.Uint32(byt)
}

func Uint32ToBytes(i uint32, byt []byte) {
	binary.LittleEndian.PutUint32(byt, i)
}
func BytesToUint64(byt []byte) uint64 {
	return binary.LittleEndian.Uint64(byt)
}

func Uint64ToBytes(i uint64, byt []byte) {
	binary.LittleEndian.PutUint64(byt, i)
}
//go:proto clear

//
// int
//

//go:proto T=Ints B=T.lower S=T.sizebits
func BytesToInt(byt []byte) int {
	return int(BytesToUint(byt))
}

func IntToBytes(i int, byt []byte) {
	UintToBytes(uint(i), byt)
}
func BytesToInt8(byt []byte) int8 {
	return int8(BytesToUint8(byt))
}

func Int8ToBytes(i int8, byt []byte) {
	Uint8ToBytes(uint8(i), byt)
}
func BytesToInt16(byt []byte) int16 {
	return int16(BytesToUint16(byt))
}

func Int16ToBytes(i int16, byt []byte) {
	Uint16ToBytes(uint16(i), byt)
}
func BytesToInt32(byt []byte) int32 {
	return int32(BytesToUint32(byt))
}

func Int32ToBytes(i int32, byt []byte) {
	Uint32ToBytes(uint32(i), byt)
}
func BytesToInt64(byt []byte) int64 {
	return int64(BytesToUint64(byt))
}

func Int64ToBytes(i int64, byt []byte) {
	Uint64ToBytes(uint64(i), byt)
}
//go:proto clear

//
// float
//

// camel case is necessary for go-proto
var float32Bits = math.Float32bits
var float32FromBits = math.Float32frombits
var float64Bits = math.Float64bits
var float64FromBits = math.Float64frombits

//go:proto T=Floats B=T.lower S=T.sizebits
func BytesToFloat32(byt []byte) float32 {
	return float32FromBits(BytesToUint32(byt))
}

func Float32ToBytes(f float32, byt []byte) {
	Uint32ToBytes(float32Bits(f), byt)
}
func BytesToFloat64(byt []byte) float64 {
	return float64FromBits(BytesToUint64(byt))
}

func Float64ToBytes(f float64, byt []byte) {
	Uint64ToBytes(float64Bits(f), byt)
}
//go:proto clear

//
// bool
//

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

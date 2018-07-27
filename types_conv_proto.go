//+build proto

package ras

import (
	"encoding/binary"
	"math"
)

//go:proto ignore
type B int

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
func BytesToT(byt []byte) B {
	return binary.LittleEndian.UintS(byt)
}

func TToBytes(i B, byt []byte) {
	binary.LittleEndian.PutUintS(byt, i)
}
//go:proto clear

//
// int
//

//go:proto T=Ints B=T.lower S=T.sizebits
func BytesToT(byt []byte) B {
	return B(BytesToUintS(byt))
}

func TToBytes(i B, byt []byte) {
	UintSToBytes(uintS(i), byt)
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
func BytesToT(byt []byte) B {
	return floatSFromBits(BytesToUintS(byt))
}

func TToBytes(f B, byt []byte) {
	UintSToBytes(floatSBits(f), byt)
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
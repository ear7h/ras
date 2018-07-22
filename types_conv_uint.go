// this file is generated
package ras


import (
	"encoding/binary"
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

func Uint16ToBytes(i uint16, byt []byte) {
	binary.LittleEndian.PutUint16(byt, i)
}

func BytesToUint16(byt []byte) uint16 {
	return binary.LittleEndian.Uint16(byt)
}

func Uint32ToBytes(i uint32, byt []byte) {
	binary.LittleEndian.PutUint32(byt, i)
}

func BytesToUint32(byt []byte) uint32 {
	return binary.LittleEndian.Uint32(byt)
}

func Uint64ToBytes(i uint64, byt []byte) {
	binary.LittleEndian.PutUint64(byt, i)
}

func BytesToUint64(byt []byte) uint64 {
	return binary.LittleEndian.Uint64(byt)
}


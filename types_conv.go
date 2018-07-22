package ras

import "math"

//
// dynamics
//

func UintToBytes(i uint, byt []byte) {
	Uint64ToBytes(uint64(i), byt)
}

func BytesToUint(byt []byte) uint {
	return uint(BytesToUint64(byt))
}

func IntToBytes(i int, byt []byte) {
	Uint64ToBytes(uint64(i), byt)
}

func BytesToInt(byt []byte) int {
	return int(BytesToUint64(byt))
}

//
// ints
//

func Int8ToBytes(i int8, byt []byte) {
	Uint8ToBytes(uint8(i), byt)
}

func BytesToInt8(byt []byte) int8 {
	return int8(BytesToUint8(byt))
}

func Int16ToBytes(i int16, byt []byte) {
	Uint16ToBytes(uint16(i), byt)
}

func BytesToInt16(byt []byte) int16 {
	return int16(BytesToUint16(byt))
}

func Int32ToBytes(i int32, byt []byte) {
	Uint32ToBytes(uint32(i), byt)
}

func BytesToInt32(byt []byte) int32 {
	return int32(BytesToUint32(byt))
}

func Int64ToBytes(i int64, byt []byte) {
	Uint64ToBytes(uint64(i), byt)
}

func BytesToInt64(byt []byte) int64 {
	return int64(BytesToUint64(byt))
}

//
// floats
//

func Float32ToBytes(f float32, byt []byte) {
	Uint32ToBytes(math.Float32bits(f), byt)
}

func BytesToFloat32(byt []byte) float32 {
	return math.Float32frombits(BytesToUint32(byt))
}

func Float64ToBytes(f float64, byt []byte) {
	Uint64ToBytes(math.Float64bits(f), byt)
}

func BytesToFloat64(byt []byte) float64 {
	return math.Float64frombits(BytesToUint64(byt))
}

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
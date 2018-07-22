
package ras

//go:proto T=UintN,IntN,Floats,Bool B=T.lower
type Uint8 struct {
	addr uint
	a Allocator
}

func (v Uint8) Addr() uint {
	return v.addr
}

func (v Uint8) SizeOf() uint {
	return Uint8Size
}

func (v Uint8) Get() (uint8, error) {
	byt := make([]byte, Uint8Size)
	err := v.a.Get(v.addr, byt)
	return BytesToUint8(byt), err

}

func (v Uint8) Set(val uint8) error {
	byt := make([]byte, Uint8Size)
	Uint8ToBytes(val, byt)
	return v.a.Set(v.addr, byt)
}
type Uint16 struct {
	addr uint
	a Allocator
}

func (v Uint16) Addr() uint {
	return v.addr
}

func (v Uint16) SizeOf() uint {
	return Uint16Size
}

func (v Uint16) Get() (uint16, error) {
	byt := make([]byte, Uint16Size)
	err := v.a.Get(v.addr, byt)
	return BytesToUint16(byt), err

}

func (v Uint16) Set(val uint16) error {
	byt := make([]byte, Uint16Size)
	Uint16ToBytes(val, byt)
	return v.a.Set(v.addr, byt)
}
type Uint32 struct {
	addr uint
	a Allocator
}

func (v Uint32) Addr() uint {
	return v.addr
}

func (v Uint32) SizeOf() uint {
	return Uint32Size
}

func (v Uint32) Get() (uint32, error) {
	byt := make([]byte, Uint32Size)
	err := v.a.Get(v.addr, byt)
	return BytesToUint32(byt), err

}

func (v Uint32) Set(val uint32) error {
	byt := make([]byte, Uint32Size)
	Uint32ToBytes(val, byt)
	return v.a.Set(v.addr, byt)
}
type Uint64 struct {
	addr uint
	a Allocator
}

func (v Uint64) Addr() uint {
	return v.addr
}

func (v Uint64) SizeOf() uint {
	return Uint64Size
}

func (v Uint64) Get() (uint64, error) {
	byt := make([]byte, Uint64Size)
	err := v.a.Get(v.addr, byt)
	return BytesToUint64(byt), err

}

func (v Uint64) Set(val uint64) error {
	byt := make([]byte, Uint64Size)
	Uint64ToBytes(val, byt)
	return v.a.Set(v.addr, byt)
}
type Int8 struct {
	addr uint
	a Allocator
}

func (v Int8) Addr() uint {
	return v.addr
}

func (v Int8) SizeOf() uint {
	return Int8Size
}

func (v Int8) Get() (int8, error) {
	byt := make([]byte, Int8Size)
	err := v.a.Get(v.addr, byt)
	return BytesToInt8(byt), err

}

func (v Int8) Set(val int8) error {
	byt := make([]byte, Int8Size)
	Int8ToBytes(val, byt)
	return v.a.Set(v.addr, byt)
}
type Int16 struct {
	addr uint
	a Allocator
}

func (v Int16) Addr() uint {
	return v.addr
}

func (v Int16) SizeOf() uint {
	return Int16Size
}

func (v Int16) Get() (int16, error) {
	byt := make([]byte, Int16Size)
	err := v.a.Get(v.addr, byt)
	return BytesToInt16(byt), err

}

func (v Int16) Set(val int16) error {
	byt := make([]byte, Int16Size)
	Int16ToBytes(val, byt)
	return v.a.Set(v.addr, byt)
}
type Int32 struct {
	addr uint
	a Allocator
}

func (v Int32) Addr() uint {
	return v.addr
}

func (v Int32) SizeOf() uint {
	return Int32Size
}

func (v Int32) Get() (int32, error) {
	byt := make([]byte, Int32Size)
	err := v.a.Get(v.addr, byt)
	return BytesToInt32(byt), err

}

func (v Int32) Set(val int32) error {
	byt := make([]byte, Int32Size)
	Int32ToBytes(val, byt)
	return v.a.Set(v.addr, byt)
}
type Int64 struct {
	addr uint
	a Allocator
}

func (v Int64) Addr() uint {
	return v.addr
}

func (v Int64) SizeOf() uint {
	return Int64Size
}

func (v Int64) Get() (int64, error) {
	byt := make([]byte, Int64Size)
	err := v.a.Get(v.addr, byt)
	return BytesToInt64(byt), err

}

func (v Int64) Set(val int64) error {
	byt := make([]byte, Int64Size)
	Int64ToBytes(val, byt)
	return v.a.Set(v.addr, byt)
}
type Float32 struct {
	addr uint
	a Allocator
}

func (v Float32) Addr() uint {
	return v.addr
}

func (v Float32) SizeOf() uint {
	return Float32Size
}

func (v Float32) Get() (float32, error) {
	byt := make([]byte, Float32Size)
	err := v.a.Get(v.addr, byt)
	return BytesToFloat32(byt), err

}

func (v Float32) Set(val float32) error {
	byt := make([]byte, Float32Size)
	Float32ToBytes(val, byt)
	return v.a.Set(v.addr, byt)
}
type Float64 struct {
	addr uint
	a Allocator
}

func (v Float64) Addr() uint {
	return v.addr
}

func (v Float64) SizeOf() uint {
	return Float64Size
}

func (v Float64) Get() (float64, error) {
	byt := make([]byte, Float64Size)
	err := v.a.Get(v.addr, byt)
	return BytesToFloat64(byt), err

}

func (v Float64) Set(val float64) error {
	byt := make([]byte, Float64Size)
	Float64ToBytes(val, byt)
	return v.a.Set(v.addr, byt)
}
type Bool struct {
	addr uint
	a Allocator
}

func (v Bool) Addr() uint {
	return v.addr
}

func (v Bool) SizeOf() uint {
	return BoolSize
}

func (v Bool) Get() (bool, error) {
	byt := make([]byte, BoolSize)
	err := v.a.Get(v.addr, byt)
	return BytesToBool(byt), err

}

func (v Bool) Set(val bool) error {
	byt := make([]byte, BoolSize)
	BoolToBytes(val, byt)
	return v.a.Set(v.addr, byt)
}
//go:proto clear

//go:proto T=Uint,Int B=T.lower
type Uint struct {
	Uint64
}

func (v Uint) Get() (uint, error) {
	i, err := v.Uint64.Get()
	return uint(i), err
}

func (v Uint) Set(val uint) error {
	return v.Uint64.Set(uint64(val))
}
type Int struct {
	Uint64
}

func (v Int) Get() (int, error) {
	i, err := v.Uint64.Get()
	return int(i), err
}

func (v Int) Set(val int) error {
	return v.Uint64.Set(uint64(val))
}


package ras


type UnsafePointer struct {
	Uint
}

func (p UnsafePointer) ptr() UnsafePointer {
	return p
}

func (p UnsafePointer) Addr() uint {
	return p.addr
}

func (_ UnsafePointer) SizeOf() uint {
	return PointerSize
}

func (p UnsafePointer) Get() (uint, error) {
	return p.Uint.Get()
}

func (p UnsafePointer) Set(val uint) error {
	return p.Uint.Set(val)
}

//go:proto T=UintN,IntN,Floats,Bool
type Uint8Pointer struct {
	UnsafePointer
}

func (p Uint8Pointer) Get() (Uint8, error) {
	addr, err := p.UnsafePointer.Get()

	return Uint8{variable{addr, p.a}}, err
}

func (p Uint8Pointer) Set(val Uint8) error {
	return p.UnsafePointer.Set(val.addr)
}
type Uint16Pointer struct {
	UnsafePointer
}

func (p Uint16Pointer) Get() (Uint16, error) {
	addr, err := p.UnsafePointer.Get()

	return Uint16{variable{addr, p.a}}, err
}

func (p Uint16Pointer) Set(val Uint16) error {
	return p.UnsafePointer.Set(val.addr)
}
type Uint32Pointer struct {
	UnsafePointer
}

func (p Uint32Pointer) Get() (Uint32, error) {
	addr, err := p.UnsafePointer.Get()

	return Uint32{variable{addr, p.a}}, err
}

func (p Uint32Pointer) Set(val Uint32) error {
	return p.UnsafePointer.Set(val.addr)
}
type Uint64Pointer struct {
	UnsafePointer
}

func (p Uint64Pointer) Get() (Uint64, error) {
	addr, err := p.UnsafePointer.Get()

	return Uint64{variable{addr, p.a}}, err
}

func (p Uint64Pointer) Set(val Uint64) error {
	return p.UnsafePointer.Set(val.addr)
}
type Int8Pointer struct {
	UnsafePointer
}

func (p Int8Pointer) Get() (Int8, error) {
	addr, err := p.UnsafePointer.Get()

	return Int8{variable{addr, p.a}}, err
}

func (p Int8Pointer) Set(val Int8) error {
	return p.UnsafePointer.Set(val.addr)
}
type Int16Pointer struct {
	UnsafePointer
}

func (p Int16Pointer) Get() (Int16, error) {
	addr, err := p.UnsafePointer.Get()

	return Int16{variable{addr, p.a}}, err
}

func (p Int16Pointer) Set(val Int16) error {
	return p.UnsafePointer.Set(val.addr)
}
type Int32Pointer struct {
	UnsafePointer
}

func (p Int32Pointer) Get() (Int32, error) {
	addr, err := p.UnsafePointer.Get()

	return Int32{variable{addr, p.a}}, err
}

func (p Int32Pointer) Set(val Int32) error {
	return p.UnsafePointer.Set(val.addr)
}
type Int64Pointer struct {
	UnsafePointer
}

func (p Int64Pointer) Get() (Int64, error) {
	addr, err := p.UnsafePointer.Get()

	return Int64{variable{addr, p.a}}, err
}

func (p Int64Pointer) Set(val Int64) error {
	return p.UnsafePointer.Set(val.addr)
}
type Float32Pointer struct {
	UnsafePointer
}

func (p Float32Pointer) Get() (Float32, error) {
	addr, err := p.UnsafePointer.Get()

	return Float32{variable{addr, p.a}}, err
}

func (p Float32Pointer) Set(val Float32) error {
	return p.UnsafePointer.Set(val.addr)
}
type Float64Pointer struct {
	UnsafePointer
}

func (p Float64Pointer) Get() (Float64, error) {
	addr, err := p.UnsafePointer.Get()

	return Float64{variable{addr, p.a}}, err
}

func (p Float64Pointer) Set(val Float64) error {
	return p.UnsafePointer.Set(val.addr)
}
type BoolPointer struct {
	UnsafePointer
}

func (p BoolPointer) Get() (Bool, error) {
	addr, err := p.UnsafePointer.Get()

	return Bool{variable{addr, p.a}}, err
}

func (p BoolPointer) Set(val Bool) error {
	return p.UnsafePointer.Set(val.addr)
}
//go:proto clear

//go:proto T=Uint,Int
type UintPointer struct {
	UnsafePointer
}

func (p UintPointer) Get() (Uint, error) {
	addr, err := p.UnsafePointer.Get()

	return Uint{Uint64{variable{addr, p.a}}}, err
}

func (p UintPointer) Set(val Uint) error {
	return p.UnsafePointer.Set(val.addr)
}
type IntPointer struct {
	UnsafePointer
}

func (p IntPointer) Get() (Int, error) {
	addr, err := p.UnsafePointer.Get()

	return Int{Uint64{variable{addr, p.a}}}, err
}

func (p IntPointer) Set(val Int) error {
	return p.UnsafePointer.Set(val.addr)
}
//go:proto clear

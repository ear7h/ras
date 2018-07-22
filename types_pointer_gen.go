
package ras




//go:proto T=Builtin
type UintPointer struct {
	addr uint
	a Allocator
}

func (p UintPointer) uint() Uint {
	return Uint{Uint64{addr: p.addr, a: p.a}}
}

func (p UintPointer) Addr() uint {
	return p.addr
}

func (p UintPointer) SizeOf() uint {
	return PointerSize
}

func (p UintPointer) Get() (uint, error) {
	return p.uint().Get()
}

func (p UintPointer) Set(val Uint) error {
	return p.uint().Set(val.addr)
}

func (p UintPointer) ToArr() (UintArray, error) {
	ptr, err := p.Get()
	if err != nil {
		return UintArray{}, err
	}

	return UintArray{
		addr: ptr,
		a: p.a,
	}, nil
}
type Uint8Pointer struct {
	addr uint
	a Allocator
}

func (p Uint8Pointer) uint() Uint {
	return Uint{Uint64{addr: p.addr, a: p.a}}
}

func (p Uint8Pointer) Addr() uint {
	return p.addr
}

func (p Uint8Pointer) SizeOf() uint {
	return PointerSize
}

func (p Uint8Pointer) Get() (uint, error) {
	return p.uint().Get()
}

func (p Uint8Pointer) Set(val Uint8) error {
	return p.uint().Set(val.addr)
}

func (p Uint8Pointer) ToArr() (Uint8Array, error) {
	ptr, err := p.Get()
	if err != nil {
		return Uint8Array{}, err
	}

	return Uint8Array{
		addr: ptr,
		a: p.a,
	}, nil
}
type Uint16Pointer struct {
	addr uint
	a Allocator
}

func (p Uint16Pointer) uint() Uint {
	return Uint{Uint64{addr: p.addr, a: p.a}}
}

func (p Uint16Pointer) Addr() uint {
	return p.addr
}

func (p Uint16Pointer) SizeOf() uint {
	return PointerSize
}

func (p Uint16Pointer) Get() (uint, error) {
	return p.uint().Get()
}

func (p Uint16Pointer) Set(val Uint16) error {
	return p.uint().Set(val.addr)
}

func (p Uint16Pointer) ToArr() (Uint16Array, error) {
	ptr, err := p.Get()
	if err != nil {
		return Uint16Array{}, err
	}

	return Uint16Array{
		addr: ptr,
		a: p.a,
	}, nil
}
type Uint32Pointer struct {
	addr uint
	a Allocator
}

func (p Uint32Pointer) uint() Uint {
	return Uint{Uint64{addr: p.addr, a: p.a}}
}

func (p Uint32Pointer) Addr() uint {
	return p.addr
}

func (p Uint32Pointer) SizeOf() uint {
	return PointerSize
}

func (p Uint32Pointer) Get() (uint, error) {
	return p.uint().Get()
}

func (p Uint32Pointer) Set(val Uint32) error {
	return p.uint().Set(val.addr)
}

func (p Uint32Pointer) ToArr() (Uint32Array, error) {
	ptr, err := p.Get()
	if err != nil {
		return Uint32Array{}, err
	}

	return Uint32Array{
		addr: ptr,
		a: p.a,
	}, nil
}
type Uint64Pointer struct {
	addr uint
	a Allocator
}

func (p Uint64Pointer) uint() Uint {
	return Uint{Uint64{addr: p.addr, a: p.a}}
}

func (p Uint64Pointer) Addr() uint {
	return p.addr
}

func (p Uint64Pointer) SizeOf() uint {
	return PointerSize
}

func (p Uint64Pointer) Get() (uint, error) {
	return p.uint().Get()
}

func (p Uint64Pointer) Set(val Uint64) error {
	return p.uint().Set(val.addr)
}

func (p Uint64Pointer) ToArr() (Uint64Array, error) {
	ptr, err := p.Get()
	if err != nil {
		return Uint64Array{}, err
	}

	return Uint64Array{
		addr: ptr,
		a: p.a,
	}, nil
}
type IntPointer struct {
	addr uint
	a Allocator
}

func (p IntPointer) uint() Uint {
	return Uint{Uint64{addr: p.addr, a: p.a}}
}

func (p IntPointer) Addr() uint {
	return p.addr
}

func (p IntPointer) SizeOf() uint {
	return PointerSize
}

func (p IntPointer) Get() (uint, error) {
	return p.uint().Get()
}

func (p IntPointer) Set(val Int) error {
	return p.uint().Set(val.addr)
}

func (p IntPointer) ToArr() (IntArray, error) {
	ptr, err := p.Get()
	if err != nil {
		return IntArray{}, err
	}

	return IntArray{
		addr: ptr,
		a: p.a,
	}, nil
}
type Int8Pointer struct {
	addr uint
	a Allocator
}

func (p Int8Pointer) uint() Uint {
	return Uint{Uint64{addr: p.addr, a: p.a}}
}

func (p Int8Pointer) Addr() uint {
	return p.addr
}

func (p Int8Pointer) SizeOf() uint {
	return PointerSize
}

func (p Int8Pointer) Get() (uint, error) {
	return p.uint().Get()
}

func (p Int8Pointer) Set(val Int8) error {
	return p.uint().Set(val.addr)
}

func (p Int8Pointer) ToArr() (Int8Array, error) {
	ptr, err := p.Get()
	if err != nil {
		return Int8Array{}, err
	}

	return Int8Array{
		addr: ptr,
		a: p.a,
	}, nil
}
type Int16Pointer struct {
	addr uint
	a Allocator
}

func (p Int16Pointer) uint() Uint {
	return Uint{Uint64{addr: p.addr, a: p.a}}
}

func (p Int16Pointer) Addr() uint {
	return p.addr
}

func (p Int16Pointer) SizeOf() uint {
	return PointerSize
}

func (p Int16Pointer) Get() (uint, error) {
	return p.uint().Get()
}

func (p Int16Pointer) Set(val Int16) error {
	return p.uint().Set(val.addr)
}

func (p Int16Pointer) ToArr() (Int16Array, error) {
	ptr, err := p.Get()
	if err != nil {
		return Int16Array{}, err
	}

	return Int16Array{
		addr: ptr,
		a: p.a,
	}, nil
}
type Int32Pointer struct {
	addr uint
	a Allocator
}

func (p Int32Pointer) uint() Uint {
	return Uint{Uint64{addr: p.addr, a: p.a}}
}

func (p Int32Pointer) Addr() uint {
	return p.addr
}

func (p Int32Pointer) SizeOf() uint {
	return PointerSize
}

func (p Int32Pointer) Get() (uint, error) {
	return p.uint().Get()
}

func (p Int32Pointer) Set(val Int32) error {
	return p.uint().Set(val.addr)
}

func (p Int32Pointer) ToArr() (Int32Array, error) {
	ptr, err := p.Get()
	if err != nil {
		return Int32Array{}, err
	}

	return Int32Array{
		addr: ptr,
		a: p.a,
	}, nil
}
type Int64Pointer struct {
	addr uint
	a Allocator
}

func (p Int64Pointer) uint() Uint {
	return Uint{Uint64{addr: p.addr, a: p.a}}
}

func (p Int64Pointer) Addr() uint {
	return p.addr
}

func (p Int64Pointer) SizeOf() uint {
	return PointerSize
}

func (p Int64Pointer) Get() (uint, error) {
	return p.uint().Get()
}

func (p Int64Pointer) Set(val Int64) error {
	return p.uint().Set(val.addr)
}

func (p Int64Pointer) ToArr() (Int64Array, error) {
	ptr, err := p.Get()
	if err != nil {
		return Int64Array{}, err
	}

	return Int64Array{
		addr: ptr,
		a: p.a,
	}, nil
}
type Float32Pointer struct {
	addr uint
	a Allocator
}

func (p Float32Pointer) uint() Uint {
	return Uint{Uint64{addr: p.addr, a: p.a}}
}

func (p Float32Pointer) Addr() uint {
	return p.addr
}

func (p Float32Pointer) SizeOf() uint {
	return PointerSize
}

func (p Float32Pointer) Get() (uint, error) {
	return p.uint().Get()
}

func (p Float32Pointer) Set(val Float32) error {
	return p.uint().Set(val.addr)
}

func (p Float32Pointer) ToArr() (Float32Array, error) {
	ptr, err := p.Get()
	if err != nil {
		return Float32Array{}, err
	}

	return Float32Array{
		addr: ptr,
		a: p.a,
	}, nil
}
type Float64Pointer struct {
	addr uint
	a Allocator
}

func (p Float64Pointer) uint() Uint {
	return Uint{Uint64{addr: p.addr, a: p.a}}
}

func (p Float64Pointer) Addr() uint {
	return p.addr
}

func (p Float64Pointer) SizeOf() uint {
	return PointerSize
}

func (p Float64Pointer) Get() (uint, error) {
	return p.uint().Get()
}

func (p Float64Pointer) Set(val Float64) error {
	return p.uint().Set(val.addr)
}

func (p Float64Pointer) ToArr() (Float64Array, error) {
	ptr, err := p.Get()
	if err != nil {
		return Float64Array{}, err
	}

	return Float64Array{
		addr: ptr,
		a: p.a,
	}, nil
}
type BoolPointer struct {
	addr uint
	a Allocator
}

func (p BoolPointer) uint() Uint {
	return Uint{Uint64{addr: p.addr, a: p.a}}
}

func (p BoolPointer) Addr() uint {
	return p.addr
}

func (p BoolPointer) SizeOf() uint {
	return PointerSize
}

func (p BoolPointer) Get() (uint, error) {
	return p.uint().Get()
}

func (p BoolPointer) Set(val Bool) error {
	return p.uint().Set(val.addr)
}

func (p BoolPointer) ToArr() (BoolArray, error) {
	ptr, err := p.Get()
	if err != nil {
		return BoolArray{}, err
	}

	return BoolArray{
		addr: ptr,
		a: p.a,
	}, nil
}

//+build proto

package ras

//go:proto ignore
type T Int

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
type TPointer struct {
	UnsafePointer
}

func (p TPointer) Get() (T, error) {
	addr, err := p.UnsafePointer.Get()

	return T{variable{addr, p.a}}, err
}

func (p TPointer) Set(val T) error {
	return p.UnsafePointer.Set(val.addr)
}
//go:proto clear

//go:proto T=Uint,Int
type TPointer struct {
	UnsafePointer
}

func (p TPointer) Get() (T, error) {
	addr, err := p.UnsafePointer.Get()

	return T{Uint64{variable{addr, p.a}}}, err
}

func (p TPointer) Set(val T) error {
	return p.UnsafePointer.Set(val.addr)
}
//go:proto clear
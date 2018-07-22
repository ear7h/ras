//+build proto

package ras

//go:proto ignore
type T Int

//go:proto ignore
type TArray IntArray


//go:proto T=Builtin
type TPointer struct {
	addr uint
	a Allocator
}

func (p TPointer) uint() Uint {
	return Uint{Uint64{addr: p.addr, a: p.a}}
}

func (p TPointer) Addr() uint {
	return p.addr
}

func (p TPointer) SizeOf() uint {
	return Uint64Size
}

func (p TPointer) Get() (uint, error) {
	return p.uint().Get()
}

func (p TPointer) Set(val T) error {
	return p.uint().Set(val.addr)
}

func (p TPointer) ToArr() (TArray, error) {
	ptr, err := p.Get()
	if err != nil {
		return TArray{}, err
	}

	return TArray{
		addr: ptr,
		a: p.a,
	}, nil
}
//go:proto T:/
//+build proto

package ras

//go:proto ignore
type B int
//go:proto ignore

//go:proto T=UintN,IntN,Floats,Bool B=T.lower
type T struct {
	addr uint
	a Allocator
}

func (v T) Addr() uint {
	return v.addr
}

func (v T) SizeOf() uint {
	return TSize
}

func (v T) Get() (B, error) {
	byt := make([]byte, TSize)
	err := v.a.Get(v.addr, byt)
	return BytesToT(byt), err

}

func (v T) Set(val B) error {
	byt := make([]byte, TSize)
	TToBytes(val, byt)
	return v.a.Set(v.addr, byt)
}
//go:proto clear

//go:proto T=Uint,Int B=T.lower
type T struct {
	Uint64
}

func (v T) Get() (B, error) {
	i, err := v.Uint64.Get()
	return B(i), err
}

func (v T) Set(val B) error {
	return v.Uint64.Set(uint64(val))
}
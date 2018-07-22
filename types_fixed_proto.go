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
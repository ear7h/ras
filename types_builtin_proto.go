//+build proto

package ras

import "errors"

//go:proto ignore
type B int

var ErrUninitializedVariable = errors.New("uninitialized variable")


//go:proto T=UintN,IntN,Floats,Bool B=T.lower
type T struct {
	variable
}

func (v T) SizeOf() uint {
	return TSize
}

func (v T) Get() (ret B, err error) {
	if v.addr == 0 {
		return ret, ErrUninitializedVariable
	}

	byt := make([]byte, TSize)
	err = v.a.Get(v.addr, byt)
	return BytesToT(byt), err

}

func (v T) Set(val B) error {
	if v.addr == 0 {
		return ErrUninitializedVariable
	}

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
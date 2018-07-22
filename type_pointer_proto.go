//+build proto

package ras

type T Int
type TArray IntArray

//go:proto kind=T
func (p Pointer) AsTArray() (TArray, error) {
	addr, err := p.Get()
	if err != nil {
		return TArray{}, err
	}

	return TArray{
		addr: addr,
		a: p.a,
	}, nil
}

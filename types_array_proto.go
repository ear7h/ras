//+build proto

package ras

//go:proto ignore
type B int
//go:proto ignore
type T Int
//go:proto ignore
const TSize = Int64Size

//go:proto T=Builtin B=T.lower
type TArray struct {
	addr uint
	a    Allocator
}

// fill dst with B's with an offset of o
func (arr TArray) Get(dst []B, o uint) (error) {
	byt := make([]byte, len(dst)*TSize)

	err := arr.a.Get((arr.addr+o)*TSize, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = BytesToT(byt[i*TSize:(i+1)*TSize])
	}

	return nil
}

// place src's B's into storage with an offset of o
func (arr TArray) Set(src []B, o uint) error {
	byt := make([]byte, len(src)*TSize)

	for i, v := range src {
		TToBytes(v, byt[i*TSize:(i+1)*TSize])
	}

	return arr.a.Set((arr.addr+o)*TSize, byt)
}


func (arr TArray)GetI(i uint)  (B, error) {
	dst := make([]B, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr TArray)SetI(val B, i uint)  error {
	src := []B{val}
	err := arr.Get(src, i)
	return err
}
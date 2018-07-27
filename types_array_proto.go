//+build proto

package ras

import "errors"

//go:proto ignore
type B int
//go:proto ignore
type T Int
//go:proto ignore
const TSize = IntSize
//go:proto ignore
type TPointer IntPointer

// When
var ErrBreakIter = errors.New("break iter")

//go:proto T=Builtin B=T.lower
type TArray struct {
	TPointer
}

func (arr TArray) SizeOfElement() uint {
	return TSize
}

// fill dst with B's with an offset of o
func (arr TArray) Get(dst []B, o uint) (error) {
	// get address from pointer
	addr, err := arr.ptr().Get()
	if err != nil {
		return err
	}

	byt := make([]byte, len(dst)*TSize)

	err = arr.a.Get((addr+o)*TSize, byt)
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
	// get address from pointer
	addr, err := arr.ptr().Get()
	if err != nil {
		return err
	}

	byt := make([]byte, len(src)*TSize)

	for i, v := range src {
		TToBytes(v, byt[i*TSize:(i+1)*TSize])
	}

	return arr.a.Set((addr+o)*TSize, byt)
}


func (arr TArray)GetI(i uint)  (B, error) {
	dst := make([]B, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr TArray)SetI(i uint, val B)  error {
	src := []B{val}
	err := arr.Get(src, i)
	return err
}

type TIterator = func(i uint, v B) error

func (arr TArray) Range(offset uint, f TIterator) error {
	var err error
	for i := uint(0);err == nil; i++ {
		v, err := arr.GetI(i)
		if err != nil {
			break
		}

		err = f(i, v)
	}

	if err == ErrBreakIter {
		return nil
	}
	return err
}
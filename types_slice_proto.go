// +build proto

package ras

//go:proto T=Builtin K=T.lower
type TSlice struct {
	variable
}

func (slc TSlice) len() Uint {
	return Uint{
		Uint64{slc.variable},
	}
}

func (slc TSlice) ptr() UnsafePointer {
	return UnsafePointer{
		Uint{
			Uint64{
				variable{
					addr: slc.addr + 8,
					a:    slc.a,
				},
			},
		},
	}
}

func (slc TSlice) arr() TArray {
	return TArray{
		TPointer{slc.ptr()},
	}
}

func (slc TSlice) Cap() (uint, error) {
	return slc.a.Bsize(slc.addr)
}

func (slc TSlice) Len() (uint, error) {
	return slc.len().Get()
}

func (slc TSlice) GetI(i uint) (v K, err error) {
	l, err := slc.Len()
	if err != nil {
		return v, err
	}
	if i >= l {
		return v, ErrSliceOutOfBounds
	}

	return slc.arr().GetI(i)
}

func (slc TSlice) SetI(i uint, v K) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	return slc.arr().SetI(i, v)
}

func (slc TSlice) Range(offset uint, f TIterator) error {
	return slc.arr().Range(offset, f)
}

func (_ TSlice) SizeOf() uint {
	return UintSize + PointerSize
}


package ras

//go:proto T=Builtin K=T.lower
type UintSlice struct {
	variable
}

func (slc UintSlice) len() Uint {
	return Uint{
		Uint64{slc.variable},
	}
}

func (slc UintSlice) ptr() UnsafePointer {
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

func (slc UintSlice) arr() UintArray {
	return UintArray{
		UintPointer{slc.ptr()},
	}
}

func (slc UintSlice) Cap() (uint, error) {
	return slc.a.Bsize(slc.addr)
}

func (slc UintSlice) Len() (uint, error) {
	return slc.len().Get()
}

func (slc UintSlice) GetI(i uint) (v uint, err error) {
	l, err := slc.Len()
	if err != nil {
		return v, err
	}
	if i >= l {
		return v, ErrSliceOutOfBounds
	}

	return slc.arr().GetI(i)
}

func (slc UintSlice) SetI(i uint, v uint) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	return slc.arr().SetI(i, v)
}

func (slc UintSlice) Range(offset uint, f UintIterator) error {
	return slc.arr().Range(offset, f)
}

func (_ UintSlice) SizeOf() uint {
	return UintSize + PointerSize
}
type Uint8Slice struct {
	variable
}

func (slc Uint8Slice) len() Uint {
	return Uint{
		Uint64{slc.variable},
	}
}

func (slc Uint8Slice) ptr() UnsafePointer {
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

func (slc Uint8Slice) arr() Uint8Array {
	return Uint8Array{
		Uint8Pointer{slc.ptr()},
	}
}

func (slc Uint8Slice) Cap() (uint, error) {
	return slc.a.Bsize(slc.addr)
}

func (slc Uint8Slice) Len() (uint, error) {
	return slc.len().Get()
}

func (slc Uint8Slice) GetI(i uint) (v uint8, err error) {
	l, err := slc.Len()
	if err != nil {
		return v, err
	}
	if i >= l {
		return v, ErrSliceOutOfBounds
	}

	return slc.arr().GetI(i)
}

func (slc Uint8Slice) SetI(i uint, v uint8) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	return slc.arr().SetI(i, v)
}

func (slc Uint8Slice) Range(offset uint, f Uint8Iterator) error {
	return slc.arr().Range(offset, f)
}

func (_ Uint8Slice) SizeOf() uint {
	return UintSize + PointerSize
}
type Uint16Slice struct {
	variable
}

func (slc Uint16Slice) len() Uint {
	return Uint{
		Uint64{slc.variable},
	}
}

func (slc Uint16Slice) ptr() UnsafePointer {
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

func (slc Uint16Slice) arr() Uint16Array {
	return Uint16Array{
		Uint16Pointer{slc.ptr()},
	}
}

func (slc Uint16Slice) Cap() (uint, error) {
	return slc.a.Bsize(slc.addr)
}

func (slc Uint16Slice) Len() (uint, error) {
	return slc.len().Get()
}

func (slc Uint16Slice) GetI(i uint) (v uint16, err error) {
	l, err := slc.Len()
	if err != nil {
		return v, err
	}
	if i >= l {
		return v, ErrSliceOutOfBounds
	}

	return slc.arr().GetI(i)
}

func (slc Uint16Slice) SetI(i uint, v uint16) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	return slc.arr().SetI(i, v)
}

func (slc Uint16Slice) Range(offset uint, f Uint16Iterator) error {
	return slc.arr().Range(offset, f)
}

func (_ Uint16Slice) SizeOf() uint {
	return UintSize + PointerSize
}
type Uint32Slice struct {
	variable
}

func (slc Uint32Slice) len() Uint {
	return Uint{
		Uint64{slc.variable},
	}
}

func (slc Uint32Slice) ptr() UnsafePointer {
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

func (slc Uint32Slice) arr() Uint32Array {
	return Uint32Array{
		Uint32Pointer{slc.ptr()},
	}
}

func (slc Uint32Slice) Cap() (uint, error) {
	return slc.a.Bsize(slc.addr)
}

func (slc Uint32Slice) Len() (uint, error) {
	return slc.len().Get()
}

func (slc Uint32Slice) GetI(i uint) (v uint32, err error) {
	l, err := slc.Len()
	if err != nil {
		return v, err
	}
	if i >= l {
		return v, ErrSliceOutOfBounds
	}

	return slc.arr().GetI(i)
}

func (slc Uint32Slice) SetI(i uint, v uint32) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	return slc.arr().SetI(i, v)
}

func (slc Uint32Slice) Range(offset uint, f Uint32Iterator) error {
	return slc.arr().Range(offset, f)
}

func (_ Uint32Slice) SizeOf() uint {
	return UintSize + PointerSize
}
type Uint64Slice struct {
	variable
}

func (slc Uint64Slice) len() Uint {
	return Uint{
		Uint64{slc.variable},
	}
}

func (slc Uint64Slice) ptr() UnsafePointer {
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

func (slc Uint64Slice) arr() Uint64Array {
	return Uint64Array{
		Uint64Pointer{slc.ptr()},
	}
}

func (slc Uint64Slice) Cap() (uint, error) {
	return slc.a.Bsize(slc.addr)
}

func (slc Uint64Slice) Len() (uint, error) {
	return slc.len().Get()
}

func (slc Uint64Slice) GetI(i uint) (v uint64, err error) {
	l, err := slc.Len()
	if err != nil {
		return v, err
	}
	if i >= l {
		return v, ErrSliceOutOfBounds
	}

	return slc.arr().GetI(i)
}

func (slc Uint64Slice) SetI(i uint, v uint64) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	return slc.arr().SetI(i, v)
}

func (slc Uint64Slice) Range(offset uint, f Uint64Iterator) error {
	return slc.arr().Range(offset, f)
}

func (_ Uint64Slice) SizeOf() uint {
	return UintSize + PointerSize
}
type IntSlice struct {
	variable
}

func (slc IntSlice) len() Uint {
	return Uint{
		Uint64{slc.variable},
	}
}

func (slc IntSlice) ptr() UnsafePointer {
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

func (slc IntSlice) arr() IntArray {
	return IntArray{
		IntPointer{slc.ptr()},
	}
}

func (slc IntSlice) Cap() (uint, error) {
	return slc.a.Bsize(slc.addr)
}

func (slc IntSlice) Len() (uint, error) {
	return slc.len().Get()
}

func (slc IntSlice) GetI(i uint) (v int, err error) {
	l, err := slc.Len()
	if err != nil {
		return v, err
	}
	if i >= l {
		return v, ErrSliceOutOfBounds
	}

	return slc.arr().GetI(i)
}

func (slc IntSlice) SetI(i uint, v int) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	return slc.arr().SetI(i, v)
}

func (slc IntSlice) Range(offset uint, f IntIterator) error {
	return slc.arr().Range(offset, f)
}

func (_ IntSlice) SizeOf() uint {
	return UintSize + PointerSize
}
type Int8Slice struct {
	variable
}

func (slc Int8Slice) len() Uint {
	return Uint{
		Uint64{slc.variable},
	}
}

func (slc Int8Slice) ptr() UnsafePointer {
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

func (slc Int8Slice) arr() Int8Array {
	return Int8Array{
		Int8Pointer{slc.ptr()},
	}
}

func (slc Int8Slice) Cap() (uint, error) {
	return slc.a.Bsize(slc.addr)
}

func (slc Int8Slice) Len() (uint, error) {
	return slc.len().Get()
}

func (slc Int8Slice) GetI(i uint) (v int8, err error) {
	l, err := slc.Len()
	if err != nil {
		return v, err
	}
	if i >= l {
		return v, ErrSliceOutOfBounds
	}

	return slc.arr().GetI(i)
}

func (slc Int8Slice) SetI(i uint, v int8) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	return slc.arr().SetI(i, v)
}

func (slc Int8Slice) Range(offset uint, f Int8Iterator) error {
	return slc.arr().Range(offset, f)
}

func (_ Int8Slice) SizeOf() uint {
	return UintSize + PointerSize
}
type Int16Slice struct {
	variable
}

func (slc Int16Slice) len() Uint {
	return Uint{
		Uint64{slc.variable},
	}
}

func (slc Int16Slice) ptr() UnsafePointer {
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

func (slc Int16Slice) arr() Int16Array {
	return Int16Array{
		Int16Pointer{slc.ptr()},
	}
}

func (slc Int16Slice) Cap() (uint, error) {
	return slc.a.Bsize(slc.addr)
}

func (slc Int16Slice) Len() (uint, error) {
	return slc.len().Get()
}

func (slc Int16Slice) GetI(i uint) (v int16, err error) {
	l, err := slc.Len()
	if err != nil {
		return v, err
	}
	if i >= l {
		return v, ErrSliceOutOfBounds
	}

	return slc.arr().GetI(i)
}

func (slc Int16Slice) SetI(i uint, v int16) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	return slc.arr().SetI(i, v)
}

func (slc Int16Slice) Range(offset uint, f Int16Iterator) error {
	return slc.arr().Range(offset, f)
}

func (_ Int16Slice) SizeOf() uint {
	return UintSize + PointerSize
}
type Int32Slice struct {
	variable
}

func (slc Int32Slice) len() Uint {
	return Uint{
		Uint64{slc.variable},
	}
}

func (slc Int32Slice) ptr() UnsafePointer {
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

func (slc Int32Slice) arr() Int32Array {
	return Int32Array{
		Int32Pointer{slc.ptr()},
	}
}

func (slc Int32Slice) Cap() (uint, error) {
	return slc.a.Bsize(slc.addr)
}

func (slc Int32Slice) Len() (uint, error) {
	return slc.len().Get()
}

func (slc Int32Slice) GetI(i uint) (v int32, err error) {
	l, err := slc.Len()
	if err != nil {
		return v, err
	}
	if i >= l {
		return v, ErrSliceOutOfBounds
	}

	return slc.arr().GetI(i)
}

func (slc Int32Slice) SetI(i uint, v int32) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	return slc.arr().SetI(i, v)
}

func (slc Int32Slice) Range(offset uint, f Int32Iterator) error {
	return slc.arr().Range(offset, f)
}

func (_ Int32Slice) SizeOf() uint {
	return UintSize + PointerSize
}
type Int64Slice struct {
	variable
}

func (slc Int64Slice) len() Uint {
	return Uint{
		Uint64{slc.variable},
	}
}

func (slc Int64Slice) ptr() UnsafePointer {
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

func (slc Int64Slice) arr() Int64Array {
	return Int64Array{
		Int64Pointer{slc.ptr()},
	}
}

func (slc Int64Slice) Cap() (uint, error) {
	return slc.a.Bsize(slc.addr)
}

func (slc Int64Slice) Len() (uint, error) {
	return slc.len().Get()
}

func (slc Int64Slice) GetI(i uint) (v int64, err error) {
	l, err := slc.Len()
	if err != nil {
		return v, err
	}
	if i >= l {
		return v, ErrSliceOutOfBounds
	}

	return slc.arr().GetI(i)
}

func (slc Int64Slice) SetI(i uint, v int64) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	return slc.arr().SetI(i, v)
}

func (slc Int64Slice) Range(offset uint, f Int64Iterator) error {
	return slc.arr().Range(offset, f)
}

func (_ Int64Slice) SizeOf() uint {
	return UintSize + PointerSize
}
type Float32Slice struct {
	variable
}

func (slc Float32Slice) len() Uint {
	return Uint{
		Uint64{slc.variable},
	}
}

func (slc Float32Slice) ptr() UnsafePointer {
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

func (slc Float32Slice) arr() Float32Array {
	return Float32Array{
		Float32Pointer{slc.ptr()},
	}
}

func (slc Float32Slice) Cap() (uint, error) {
	return slc.a.Bsize(slc.addr)
}

func (slc Float32Slice) Len() (uint, error) {
	return slc.len().Get()
}

func (slc Float32Slice) GetI(i uint) (v float32, err error) {
	l, err := slc.Len()
	if err != nil {
		return v, err
	}
	if i >= l {
		return v, ErrSliceOutOfBounds
	}

	return slc.arr().GetI(i)
}

func (slc Float32Slice) SetI(i uint, v float32) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	return slc.arr().SetI(i, v)
}

func (slc Float32Slice) Range(offset uint, f Float32Iterator) error {
	return slc.arr().Range(offset, f)
}

func (_ Float32Slice) SizeOf() uint {
	return UintSize + PointerSize
}
type Float64Slice struct {
	variable
}

func (slc Float64Slice) len() Uint {
	return Uint{
		Uint64{slc.variable},
	}
}

func (slc Float64Slice) ptr() UnsafePointer {
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

func (slc Float64Slice) arr() Float64Array {
	return Float64Array{
		Float64Pointer{slc.ptr()},
	}
}

func (slc Float64Slice) Cap() (uint, error) {
	return slc.a.Bsize(slc.addr)
}

func (slc Float64Slice) Len() (uint, error) {
	return slc.len().Get()
}

func (slc Float64Slice) GetI(i uint) (v float64, err error) {
	l, err := slc.Len()
	if err != nil {
		return v, err
	}
	if i >= l {
		return v, ErrSliceOutOfBounds
	}

	return slc.arr().GetI(i)
}

func (slc Float64Slice) SetI(i uint, v float64) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	return slc.arr().SetI(i, v)
}

func (slc Float64Slice) Range(offset uint, f Float64Iterator) error {
	return slc.arr().Range(offset, f)
}

func (_ Float64Slice) SizeOf() uint {
	return UintSize + PointerSize
}
type BoolSlice struct {
	variable
}

func (slc BoolSlice) len() Uint {
	return Uint{
		Uint64{slc.variable},
	}
}

func (slc BoolSlice) ptr() UnsafePointer {
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

func (slc BoolSlice) arr() BoolArray {
	return BoolArray{
		BoolPointer{slc.ptr()},
	}
}

func (slc BoolSlice) Cap() (uint, error) {
	return slc.a.Bsize(slc.addr)
}

func (slc BoolSlice) Len() (uint, error) {
	return slc.len().Get()
}

func (slc BoolSlice) GetI(i uint) (v bool, err error) {
	l, err := slc.Len()
	if err != nil {
		return v, err
	}
	if i >= l {
		return v, ErrSliceOutOfBounds
	}

	return slc.arr().GetI(i)
}

func (slc BoolSlice) SetI(i uint, v bool) error {
	l, err := slc.Len()
	if err != nil {
		return err
	}
	if i >= l {
		return ErrSliceOutOfBounds
	}

	return slc.arr().SetI(i, v)
}

func (slc BoolSlice) Range(offset uint, f BoolIterator) error {
	return slc.arr().Range(offset, f)
}

func (_ BoolSlice) SizeOf() uint {
	return UintSize + PointerSize
}

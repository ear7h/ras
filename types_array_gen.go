
package ras


//go:proto T=Builtin B=T.lower
type UintArray struct {
	addr uint
	a    Allocator
}

// fill dst with uint's with an offset of o
func (arr UintArray) Get(dst []uint, o uint) (error) {
	byt := make([]byte, len(dst)*UintSize)

	err := arr.a.Get((arr.addr+o)*UintSize, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = BytesToUint(byt[i*UintSize:(i+1)*UintSize])
	}

	return nil
}

// place src's uint's into storage with an offset of o
func (arr UintArray) Set(src []uint, o uint) error {
	byt := make([]byte, len(src)*UintSize)

	for i, v := range src {
		UintToBytes(v, byt[i*UintSize:(i+1)*UintSize])
	}

	return arr.a.Set((arr.addr+o)*UintSize, byt)
}


func (arr UintArray)GetI(i uint)  (uint, error) {
	dst := make([]uint, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr UintArray)SetI(val uint, i uint)  error {
	src := []uint{val}
	err := arr.Get(src, i)
	return err
}
type Uint8Array struct {
	addr uint
	a    Allocator
}

// fill dst with uint8's with an offset of o
func (arr Uint8Array) Get(dst []uint8, o uint) (error) {
	byt := make([]byte, len(dst)*Uint8Size)

	err := arr.a.Get((arr.addr+o)*Uint8Size, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = BytesToUint8(byt[i*Uint8Size:(i+1)*Uint8Size])
	}

	return nil
}

// place src's uint8's into storage with an offset of o
func (arr Uint8Array) Set(src []uint8, o uint) error {
	byt := make([]byte, len(src)*Uint8Size)

	for i, v := range src {
		Uint8ToBytes(v, byt[i*Uint8Size:(i+1)*Uint8Size])
	}

	return arr.a.Set((arr.addr+o)*Uint8Size, byt)
}


func (arr Uint8Array)GetI(i uint)  (uint8, error) {
	dst := make([]uint8, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr Uint8Array)SetI(val uint8, i uint)  error {
	src := []uint8{val}
	err := arr.Get(src, i)
	return err
}
type Uint16Array struct {
	addr uint
	a    Allocator
}

// fill dst with uint16's with an offset of o
func (arr Uint16Array) Get(dst []uint16, o uint) (error) {
	byt := make([]byte, len(dst)*Uint16Size)

	err := arr.a.Get((arr.addr+o)*Uint16Size, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = BytesToUint16(byt[i*Uint16Size:(i+1)*Uint16Size])
	}

	return nil
}

// place src's uint16's into storage with an offset of o
func (arr Uint16Array) Set(src []uint16, o uint) error {
	byt := make([]byte, len(src)*Uint16Size)

	for i, v := range src {
		Uint16ToBytes(v, byt[i*Uint16Size:(i+1)*Uint16Size])
	}

	return arr.a.Set((arr.addr+o)*Uint16Size, byt)
}


func (arr Uint16Array)GetI(i uint)  (uint16, error) {
	dst := make([]uint16, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr Uint16Array)SetI(val uint16, i uint)  error {
	src := []uint16{val}
	err := arr.Get(src, i)
	return err
}
type Uint32Array struct {
	addr uint
	a    Allocator
}

// fill dst with uint32's with an offset of o
func (arr Uint32Array) Get(dst []uint32, o uint) (error) {
	byt := make([]byte, len(dst)*Uint32Size)

	err := arr.a.Get((arr.addr+o)*Uint32Size, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = BytesToUint32(byt[i*Uint32Size:(i+1)*Uint32Size])
	}

	return nil
}

// place src's uint32's into storage with an offset of o
func (arr Uint32Array) Set(src []uint32, o uint) error {
	byt := make([]byte, len(src)*Uint32Size)

	for i, v := range src {
		Uint32ToBytes(v, byt[i*Uint32Size:(i+1)*Uint32Size])
	}

	return arr.a.Set((arr.addr+o)*Uint32Size, byt)
}


func (arr Uint32Array)GetI(i uint)  (uint32, error) {
	dst := make([]uint32, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr Uint32Array)SetI(val uint32, i uint)  error {
	src := []uint32{val}
	err := arr.Get(src, i)
	return err
}
type Uint64Array struct {
	addr uint
	a    Allocator
}

// fill dst with uint64's with an offset of o
func (arr Uint64Array) Get(dst []uint64, o uint) (error) {
	byt := make([]byte, len(dst)*Uint64Size)

	err := arr.a.Get((arr.addr+o)*Uint64Size, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = BytesToUint64(byt[i*Uint64Size:(i+1)*Uint64Size])
	}

	return nil
}

// place src's uint64's into storage with an offset of o
func (arr Uint64Array) Set(src []uint64, o uint) error {
	byt := make([]byte, len(src)*Uint64Size)

	for i, v := range src {
		Uint64ToBytes(v, byt[i*Uint64Size:(i+1)*Uint64Size])
	}

	return arr.a.Set((arr.addr+o)*Uint64Size, byt)
}


func (arr Uint64Array)GetI(i uint)  (uint64, error) {
	dst := make([]uint64, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr Uint64Array)SetI(val uint64, i uint)  error {
	src := []uint64{val}
	err := arr.Get(src, i)
	return err
}
type IntArray struct {
	addr uint
	a    Allocator
}

// fill dst with int's with an offset of o
func (arr IntArray) Get(dst []int, o uint) (error) {
	byt := make([]byte, len(dst)*IntSize)

	err := arr.a.Get((arr.addr+o)*IntSize, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = BytesToInt(byt[i*IntSize:(i+1)*IntSize])
	}

	return nil
}

// place src's int's into storage with an offset of o
func (arr IntArray) Set(src []int, o uint) error {
	byt := make([]byte, len(src)*IntSize)

	for i, v := range src {
		IntToBytes(v, byt[i*IntSize:(i+1)*IntSize])
	}

	return arr.a.Set((arr.addr+o)*IntSize, byt)
}


func (arr IntArray)GetI(i uint)  (int, error) {
	dst := make([]int, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr IntArray)SetI(val int, i uint)  error {
	src := []int{val}
	err := arr.Get(src, i)
	return err
}
type Int8Array struct {
	addr uint
	a    Allocator
}

// fill dst with int8's with an offset of o
func (arr Int8Array) Get(dst []int8, o uint) (error) {
	byt := make([]byte, len(dst)*Int8Size)

	err := arr.a.Get((arr.addr+o)*Int8Size, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = BytesToInt8(byt[i*Int8Size:(i+1)*Int8Size])
	}

	return nil
}

// place src's int8's into storage with an offset of o
func (arr Int8Array) Set(src []int8, o uint) error {
	byt := make([]byte, len(src)*Int8Size)

	for i, v := range src {
		Int8ToBytes(v, byt[i*Int8Size:(i+1)*Int8Size])
	}

	return arr.a.Set((arr.addr+o)*Int8Size, byt)
}


func (arr Int8Array)GetI(i uint)  (int8, error) {
	dst := make([]int8, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr Int8Array)SetI(val int8, i uint)  error {
	src := []int8{val}
	err := arr.Get(src, i)
	return err
}
type Int16Array struct {
	addr uint
	a    Allocator
}

// fill dst with int16's with an offset of o
func (arr Int16Array) Get(dst []int16, o uint) (error) {
	byt := make([]byte, len(dst)*Int16Size)

	err := arr.a.Get((arr.addr+o)*Int16Size, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = BytesToInt16(byt[i*Int16Size:(i+1)*Int16Size])
	}

	return nil
}

// place src's int16's into storage with an offset of o
func (arr Int16Array) Set(src []int16, o uint) error {
	byt := make([]byte, len(src)*Int16Size)

	for i, v := range src {
		Int16ToBytes(v, byt[i*Int16Size:(i+1)*Int16Size])
	}

	return arr.a.Set((arr.addr+o)*Int16Size, byt)
}


func (arr Int16Array)GetI(i uint)  (int16, error) {
	dst := make([]int16, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr Int16Array)SetI(val int16, i uint)  error {
	src := []int16{val}
	err := arr.Get(src, i)
	return err
}
type Int32Array struct {
	addr uint
	a    Allocator
}

// fill dst with int32's with an offset of o
func (arr Int32Array) Get(dst []int32, o uint) (error) {
	byt := make([]byte, len(dst)*Int32Size)

	err := arr.a.Get((arr.addr+o)*Int32Size, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = BytesToInt32(byt[i*Int32Size:(i+1)*Int32Size])
	}

	return nil
}

// place src's int32's into storage with an offset of o
func (arr Int32Array) Set(src []int32, o uint) error {
	byt := make([]byte, len(src)*Int32Size)

	for i, v := range src {
		Int32ToBytes(v, byt[i*Int32Size:(i+1)*Int32Size])
	}

	return arr.a.Set((arr.addr+o)*Int32Size, byt)
}


func (arr Int32Array)GetI(i uint)  (int32, error) {
	dst := make([]int32, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr Int32Array)SetI(val int32, i uint)  error {
	src := []int32{val}
	err := arr.Get(src, i)
	return err
}
type Int64Array struct {
	addr uint
	a    Allocator
}

// fill dst with int64's with an offset of o
func (arr Int64Array) Get(dst []int64, o uint) (error) {
	byt := make([]byte, len(dst)*Int64Size)

	err := arr.a.Get((arr.addr+o)*Int64Size, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = BytesToInt64(byt[i*Int64Size:(i+1)*Int64Size])
	}

	return nil
}

// place src's int64's into storage with an offset of o
func (arr Int64Array) Set(src []int64, o uint) error {
	byt := make([]byte, len(src)*Int64Size)

	for i, v := range src {
		Int64ToBytes(v, byt[i*Int64Size:(i+1)*Int64Size])
	}

	return arr.a.Set((arr.addr+o)*Int64Size, byt)
}


func (arr Int64Array)GetI(i uint)  (int64, error) {
	dst := make([]int64, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr Int64Array)SetI(val int64, i uint)  error {
	src := []int64{val}
	err := arr.Get(src, i)
	return err
}
type Float32Array struct {
	addr uint
	a    Allocator
}

// fill dst with float32's with an offset of o
func (arr Float32Array) Get(dst []float32, o uint) (error) {
	byt := make([]byte, len(dst)*Float32Size)

	err := arr.a.Get((arr.addr+o)*Float32Size, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = BytesToFloat32(byt[i*Float32Size:(i+1)*Float32Size])
	}

	return nil
}

// place src's float32's into storage with an offset of o
func (arr Float32Array) Set(src []float32, o uint) error {
	byt := make([]byte, len(src)*Float32Size)

	for i, v := range src {
		Float32ToBytes(v, byt[i*Float32Size:(i+1)*Float32Size])
	}

	return arr.a.Set((arr.addr+o)*Float32Size, byt)
}


func (arr Float32Array)GetI(i uint)  (float32, error) {
	dst := make([]float32, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr Float32Array)SetI(val float32, i uint)  error {
	src := []float32{val}
	err := arr.Get(src, i)
	return err
}
type Float64Array struct {
	addr uint
	a    Allocator
}

// fill dst with float64's with an offset of o
func (arr Float64Array) Get(dst []float64, o uint) (error) {
	byt := make([]byte, len(dst)*Float64Size)

	err := arr.a.Get((arr.addr+o)*Float64Size, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = BytesToFloat64(byt[i*Float64Size:(i+1)*Float64Size])
	}

	return nil
}

// place src's float64's into storage with an offset of o
func (arr Float64Array) Set(src []float64, o uint) error {
	byt := make([]byte, len(src)*Float64Size)

	for i, v := range src {
		Float64ToBytes(v, byt[i*Float64Size:(i+1)*Float64Size])
	}

	return arr.a.Set((arr.addr+o)*Float64Size, byt)
}


func (arr Float64Array)GetI(i uint)  (float64, error) {
	dst := make([]float64, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr Float64Array)SetI(val float64, i uint)  error {
	src := []float64{val}
	err := arr.Get(src, i)
	return err
}
type BoolArray struct {
	addr uint
	a    Allocator
}

// fill dst with bool's with an offset of o
func (arr BoolArray) Get(dst []bool, o uint) (error) {
	byt := make([]byte, len(dst)*BoolSize)

	err := arr.a.Get((arr.addr+o)*BoolSize, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = BytesToBool(byt[i*BoolSize:(i+1)*BoolSize])
	}

	return nil
}

// place src's bool's into storage with an offset of o
func (arr BoolArray) Set(src []bool, o uint) error {
	byt := make([]byte, len(src)*BoolSize)

	for i, v := range src {
		BoolToBytes(v, byt[i*BoolSize:(i+1)*BoolSize])
	}

	return arr.a.Set((arr.addr+o)*BoolSize, byt)
}


func (arr BoolArray)GetI(i uint)  (bool, error) {
	dst := make([]bool, 1)
	err := arr.Get(dst, i)
	return dst[0], err
}

func (arr BoolArray)SetI(val bool, i uint)  error {
	src := []bool{val}
	err := arr.Get(src, i)
	return err
}

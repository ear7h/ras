package types_gen

import (
	"reflect"
	"fmt"
	"text/template"
	"strings"
)

type Array struct {
	Kind reflect.Kind
	Get string
	Set string
}

func (a Array) TLower() string {
	return a.Kind.String()
}

func (a Array) TUpper() string {
	return strings.Title(a.TLower())+"Array"
}

const arrTmplStr = `type {{ .TUpper }} struct {
	addr uint
	a Allocator
}

func (arr {{ .TUpper }}) Get(dst []{{ .TLower }}, o uint) error {
	{{ .Get }}
}

func (arr {{ .TUpper }}) Set(src []{{ .TLower }}, o uint) error {
	{{ .Set }}
}
`

var arrTmpl = template.Must(template.New("arrTmpl").Parse(arrTmplStr))

func (a Array) String() string {
	ret := &strings.Builder{}
	ret.Grow(50)
	err := arrTmpl.Execute(ret, a)
	if err != nil {
		panic(err)
	}
	return ret.String()
}

func UintArrGet(bitSize uint) string {
	if bitSize == 0 {
		return `byt := make([]byte, len(dst) * 8)

	err := arr.a.Get(arr.addr + o*8, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = uint(binary.
			LittleEndian.
			Uint64(byt[i*8:(i+1)*8]))
	}

	return nil`
	}

	if bitSize == 8 {
		return `byt := make([]byte, len(dst))

	err := arr.a.Get(arr.addr, byt)
	if err != nil {
		return err
	}

	for i, v := range byt {
		dst[i] = uint8(v)
	}

	return nil`
	}

	fmtStr := `byt := make([]byte, len(dst) * %d)

	err := arr.a.Get(arr.addr + o*%d, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = binary.
			LittleEndian.
			Uint%d(byt[i*%d:(i+1)*%d])
	}

	return nil`

	byteSize := bitSize /8
	return fmt.Sprintf(fmtStr, byteSize, byteSize, bitSize, byteSize, byteSize)
}

func UintArrSet(bitSize uint) string {
	if bitSize == 0 {
		return `byt := make([]byte, len(src) * 8)
	for i, v := range src {
		binary.
			LittleEndian.
			PutUint64(byt[i*8:(i+1)*8], uint64(v))
	}

	return arr.a.Set(arr.addr + o*8, byt)`
	}

	if bitSize == 8 {
		return `byt := make([]byte, len(src))
	for i, v := range src {
		byt[i] = byte(v)
	}

	return arr.a.Set(arr.addr + o, byt)`
	}

	fmtStr := `byt := make([]byte, len(src) * %d)
	for i, v := range src {
		binary.
			LittleEndian.
			PutUint%d(byt[i*%d:(i+1)*%d], v)
	}

	return arr.a.Set(arr.addr + o*%d, byt)
`

	byteSize := bitSize /8
	return fmt.Sprintf(fmtStr, byteSize, bitSize, byteSize, byteSize, byteSize)
}

func IntArrGet(bitSize uint) string {
	if bitSize == 0 {
		return `byt := make([]byte, len(dst) * 8)

	err := arr.a.Get(arr.addr + o*8, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = int(binary.
			LittleEndian.
			Uint64(byt[i*8:(i+1)*8]))
	}

	return nil`
	}

	if bitSize == 8 {
		return `byt := make([]byte, len(dst))

	err := arr.a.Get(arr.addr, byt)
	if err != nil {
		return err
	}

	for i, v := range byt {
		dst[i] = int8(v)
	}

	return nil`
	}

	fmtStr := `byt := make([]byte, len(dst) * %d)

	err := arr.a.Get(arr.addr + o*%d, byt)
	if err != nil {
		return err
	}

	for i := range dst {
		dst[i] = int%d(binary.
			LittleEndian.
			Uint%d(byt[i*%d:(i+1)*%d]))
	}

	return nil`

	byteSize := bitSize /8
	return fmt.Sprintf(fmtStr, byteSize, byteSize, bitSize, bitSize, byteSize, byteSize)
}

func IntArrSet(bitSize uint) string {
	if bitSize == 0 {
		return `byt := make([]byte, len(src) * 8)
	for i, v := range src {
		binary.
			LittleEndian.
			PutUint64(byt[i*8:(i+1)*8], uint64(v))
	}

	return arr.a.Set(arr.addr + o*8, byt)`
	}

	if bitSize == 8 {
		return `byt := make([]byte, len(src))
	for i, v := range src {
		byt[i] = byte(v)
	}

	return arr.a.Set(arr.addr + o, byt)`
	}

	fmtStr := `byt := make([]byte, len(src) * %d)
	for i, v := range src {
		binary.
			LittleEndian.
			PutUint%d(byt[i*%d:(i+1)*%d], uint%d(v))
	}

	return arr.a.Set(arr.addr + o*%d, byt)`

	byteSize := bitSize /8
	return fmt.Sprintf(fmtStr, byteSize, bitSize, byteSize, byteSize, bitSize, byteSize)
}
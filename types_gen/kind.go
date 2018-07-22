package types_gen

import (
	"strings"
	"reflect"
	"text/template"
)

type Kind reflect.Kind

func (k Kind) TLower() string {
	return reflect.Kind(k).String()
}

func (k Kind) TUpper() string {
	return strings.Title(k.TLower())
}

const kindTmplStr = `type {{ .TUpper }} struct {
	addr uint
	a Allocator
}

func (v {{ .TUpper }}) Addr() uint {
	return v.addr
}

func (v {{ .TUpper }}) SizeOf() uint {
	return {{ .TUpper }}Size
}

func (v {{ .TUpper }}) Get() ({{ .TLower }}, error) {
	byt := make([]byte, {{ .TUpper }}Size)
	err := v.a.Get(v.addr, byt)
	return BytesTo{{ .TUpper }}(byt), err
	
}

func (v {{ .TUpper }}) Set(val {{ .TLower }}) error {
	byt := make([]byte, {{ .TUpper }}Size)
	{{ .TUpper }}ToBytes(val, byt)
	return v.a.Set(v.addr, byt)
}
`

var kindTmpl = template.Must(template.New("kindTmpl").Parse(kindTmplStr))

func (k Kind) String() string {
	ret := &strings.Builder{}
	ret.Grow(50)
	err := kindTmpl.Execute(ret, k)
	if err != nil {
		panic(err)
	}
	return ret.String()
}

var Uints = []reflect.Kind{
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
}

var Ints = []reflect.Kind{
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
}

var Floats = []reflect.Kind {
	reflect.Float32,
	reflect.Float64,
}
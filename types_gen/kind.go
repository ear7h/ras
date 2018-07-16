package types_gen

import (
	"strings"
	"reflect"
	"text/template"
)

type Kind struct {
	Kind reflect.Kind
	Set  string
	Get  string
}

func (t Kind) TLower() string {
	return t.Kind.String()
}

func (t Kind) TUpper() string {
	return strings.Title(t.TLower())
}

const tTmplStr = `type {{ .TUpper }} struct {
	addr uint
	a Allocator
}

func (v {{ .TUpper }}) Get() ({{ .TLower }}, error) {
	{{ .Get }}
}

func (v {{ .TUpper }}) Set(val {{ .TLower }}) error {
	{{ .Set }}
}
`

var tTmpl = template.Must(template.New("tTmpl").Parse(tTmplStr))

func (t Kind) String() string {
	ret := &strings.Builder{}
	ret.Grow(50)
	err := tTmpl.Execute(ret, t)
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
package types_gen

import (
	"text/template"
	"strings"
)

var convUintTmplStr = `func {{ .TUpper }}ToBytes(i {{ .TLower }}, byt []byte) {
	binary.LittleEndian.Put{{ .TUpper }}(byt, i)
}

func BytesTo{{ .TUpper }}(byt []byte) {{ .TLower }} {
	return binary.LittleEndian.{{ .TUpper }}(byt)
}
`

var convUintTmpl = template.Must(template.New("convUintTmpl").Parse(convUintTmplStr))

func ConvUint(k Kind) string {
	str := &strings.Builder{}
	convUintTmpl.Execute(str, k)

	return str.String()
}
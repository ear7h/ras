
package ras

const (
	//go:proto T=Bool,IntN,UintN,Floats S=T.sizebytes
	BoolSize = 1
	Int8Size = 1
	Int16Size = 2
	Int32Size = 4
	Int64Size = 8
	Uint8Size = 1
	Uint16Size = 2
	Uint32Size = 4
	Uint64Size = 8
	Float32Size = 4
	Float64Size = 8
	//go:proto clear
	UintSize    = Uint64Size
	IntSize     = Int64Size
	PointerSize = Uint64Size
)

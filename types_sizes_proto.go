//+build proto

package ras

const (
	//go:proto T=Bool,IntN,UintN,Floats S=T.sizebytes
	TSize = S
	//go:proto clear
	UintSize    = Uint64Size
	IntSize     = Int64Size
	PointerSize = Uint64Size
)

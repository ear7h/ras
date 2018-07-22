// +build none

package ras

// some builtin type
type t int

type T interface{
	SizeOf() uint
	Addr() uint

	Get() (t, error)
	Set(v t) error
}

type TArray interface {
	// returns size of pointer (8)
	SizeOf() uint
	Addr() uint

	Get(offset uint, dst []t) (error)
	Set(offset uint, src []t) error

	GetI(i uint) (t, error)
	SetI(i uint, v t) error

	Range(offset uint, func(i uint, v t) error) error
}

// The in-storage representation of an array
type TArrayProto struct {
	arr uint64 // pointer to proper memory block
}

type TSlice interface {
	TArray

	Len(uint)
	Append(v T) error
}

// The in-storage representation of a slice
type TSliceProto struct {
	len uint64 // semantically similar to a golang array
	arr uint64 // pointer to an array
}

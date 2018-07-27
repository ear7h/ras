// +build none

package ras

import "errors"

// some builtin type
type t int

type T interface{
	SizeOf() uint
	Addr() uint

	Get() (t, error)
	Set(v t) error
}

// note that arrays wrap the TPointer type
// therefore getting the pointer is done through (TArray).TPointer.Get()
type TArray interface {
	SizeOf() uint
	Addr() uint

	Get(offset uint, dst []t) (error)
	Set(offset uint, src []t) error

	GetI(i uint) (t, error)
	SetI(i uint, v t) error

	// if f returns an error, iteration is stopped
	// and the error is also returned by Range, unless
	// the error is ErrStopIter
	Range(offset uint, f TIterator) error
}

type TIterator = func(i uint, k t) error
var ErrStopIter = errors.New("stop iter")

// The in-storage representation of an array
type TArrayProto struct {
	arr uint64 // pointer to proper memory block
}

type TSlice interface {
	TArray

	Len() uint
	Append(v T) error
}

// The in-storage representation of a slice
type TSliceProto struct {
	len uint64 // semantically similar to a golang array
	arr uint64 // pointer to an array
}

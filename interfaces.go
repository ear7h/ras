package ras

type Addrer interface {
	Addr() uint
}

// Fixed size variables should be capable of describing their size
type Variable interface {
	// The address of the variable
	Addr() uint
	// Number of bytes of statically allocated, contigious storage
	// similar to sizeof in c (ie. the size of an array
	// is the same as a pointer)
	SizeOf() uint
	// used for unmarshaling structs
	init(uint, Allocator)
}

type Pointer interface{
	Variable
	ptr() UnsafePointer
}

type Array interface {
	Pointer
	SizeOfElement() uint
}

type Slice interface{
	Array
	len() Uint
	Len() (uint, error)
	Cap() (uint, error)
}

type variable struct {
	addr uint
	a Allocator
}

func (v variable) Addr() uint {
	return v.addr
}

func (v *variable) init(addr uint, a Allocator) {
	v.addr = addr
	v.a = a
}
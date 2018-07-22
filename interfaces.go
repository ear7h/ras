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
}

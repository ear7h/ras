package ras

type Allocator interface {
	Cap() uint
	Resize(uint) error

	Alloc(size uint) (addr uint, err error)
	Free(addr uint) (err error)

	Get(addr uint, dst []byte) (err error)
	Set(addr uint, data []byte) (err error)
}

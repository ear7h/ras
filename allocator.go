package ras

type Allocator interface {
	// The size of the underlying file
	Cap() uint
	// Change the size of the underlying file
	Resize(uint) error

	// Allocate at least size bytes, Returns address
	// to the beginning of usable storage.
	Alloc(size uint) (addr uint, err error)
	// Free storage associated with this address.
	// Behaviour is ambiguous when passing an address
	// not obtained through a call to Alloc
	Free(addr uint) (err error)

	// Size of usable storage (excluding possible block headers or
	// metadata) associated with the address. Behaviour is ambiguous
	// when passing an address not obtained through a call to Alloc
	Bsize(addr uint) (size uint, err error)


	// fill the dst slice with bytes starting from addr
	Get(addr uint, dst []byte) (err error)
	// write the bytes from data into the underlying file
	// starting at addr
	Set(addr uint, data []byte) (err error)
}

package ras

type Addrer interface {
	Addr() uint
}

type Runtime struct {
	Allocator
}

func (r Runtime) NewUint8() (ret Uint8, err error) {
	addr, err := r.Alloc(1)
	if err != nil {
		return ret, err
	}

	ret.addr = addr
	ret.a = r

	return ret, err
}

func (r Runtime) DerefUint8(addr uint) (Uint8, error) {
	if addr >= r.Cap() {
		return Uint8{}, ErrSegFault
	}

	return Uint8{
		addr: addr,
		a: r,
	}, nil
}
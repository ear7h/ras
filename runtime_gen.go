
package ras


//go:proto T=Builtin
func (r Runtime) NewUint() (ret Uint, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

func (r Runtime) Uint(addr uint) (ret Uint) {
	ret.init(addr, r)
	return
}
func (r Runtime) NewUint8() (ret Uint8, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

func (r Runtime) Uint8(addr uint) (ret Uint8) {
	ret.init(addr, r)
	return
}
func (r Runtime) NewUint16() (ret Uint16, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

func (r Runtime) Uint16(addr uint) (ret Uint16) {
	ret.init(addr, r)
	return
}
func (r Runtime) NewUint32() (ret Uint32, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

func (r Runtime) Uint32(addr uint) (ret Uint32) {
	ret.init(addr, r)
	return
}
func (r Runtime) NewUint64() (ret Uint64, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

func (r Runtime) Uint64(addr uint) (ret Uint64) {
	ret.init(addr, r)
	return
}
func (r Runtime) NewInt() (ret Int, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

func (r Runtime) Int(addr uint) (ret Int) {
	ret.init(addr, r)
	return
}
func (r Runtime) NewInt8() (ret Int8, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

func (r Runtime) Int8(addr uint) (ret Int8) {
	ret.init(addr, r)
	return
}
func (r Runtime) NewInt16() (ret Int16, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

func (r Runtime) Int16(addr uint) (ret Int16) {
	ret.init(addr, r)
	return
}
func (r Runtime) NewInt32() (ret Int32, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

func (r Runtime) Int32(addr uint) (ret Int32) {
	ret.init(addr, r)
	return
}
func (r Runtime) NewInt64() (ret Int64, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

func (r Runtime) Int64(addr uint) (ret Int64) {
	ret.init(addr, r)
	return
}
func (r Runtime) NewFloat32() (ret Float32, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

func (r Runtime) Float32(addr uint) (ret Float32) {
	ret.init(addr, r)
	return
}
func (r Runtime) NewFloat64() (ret Float64, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

func (r Runtime) Float64(addr uint) (ret Float64) {
	ret.init(addr, r)
	return
}
func (r Runtime) NewBool() (ret Bool, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

func (r Runtime) Bool(addr uint) (ret Bool) {
	ret.init(addr, r)
	return
}

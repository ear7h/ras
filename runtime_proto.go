//+build proto

package ras

//go:proto ignore
type T Int

//go:proto kind=T
func (r Runtime) NewT() (ret T, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

//go:proto kind=T
func (r Runtime) T(addr uint) (ret T) {
	ret.a = r
	ret.addr = addr
	return
}
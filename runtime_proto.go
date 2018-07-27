//+build proto

package ras

//go:proto ignore
type T Int

//go:proto T=Builtin
func (r Runtime) NewT() (ret T, err error) {
	ret.a = r
	ret.addr, err = r.Alloc(ret.SizeOf())
	return
}

func (r Runtime) T(addr uint) (ret T) {
	ret.init(addr, r)
	return
}
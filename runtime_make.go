package ras

func (r Runtime) Make(slc Slice, length uint) error {
	size := slc.SizeOfElement() * length
	addr, err := r.Alloc(size)
	if err != nil {
		return err
	}

	err = slc.ptr().Set(addr)
	if err != nil {
		return err
	}

	return slc.len().Set(length)
}

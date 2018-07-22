package ras

type Int struct {
	Int64
}

func (v Int) Get() (int, error) {
	i, err := v.Int64.Get()
	return int(i), err
}

func (v Int) Set(val int) error {
	return v.Int64.Set(int64(val))
}

type Uint struct {
	Uint64
}

func (v Uint) Get() (uint, error) {
	i, err := v.Uint64.Get()
	return uint(i), err
}

func (v Uint) Set(val uint) error {
	return v.Uint64.Set(uint64(val))
}

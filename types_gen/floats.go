package types_gen

import (
	"fmt"
)

func FloatSet(bitSize uint) string {
	fmtStr := `data := make([]byte, %d)
	ui := math.Float%dbits(val)
	binary.LittleEndian.PutUint%d(data, ui)
	return v.a.Set(v.addr, data)`

	return fmt.Sprintf(fmtStr, bitSize/8, bitSize, bitSize)
}

func FloatGet(bitSize uint) string {
	fmtStr := `dst := make([]byte, %d)
	err := v.a.Get(v.addr, dst)
	ui := binary.LittleEndian.Uint%d(dst)
	return math.Float%dfrombits(ui), err`

	return fmt.Sprintf(fmtStr, bitSize/8, bitSize, bitSize)
}
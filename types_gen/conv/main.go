package main

import (
	"fmt"

	tg "github.com/ear7h/ras/types_gen"
)

func main() {
	fmt.Print("// this file is generated\n")
	fmt.Print("package ras\n\n")
	fmt.Print(`
import (
	"encoding/binary"
)

`)

	fmt.Print(`//
// uints
//
`)

	fmt.Print(`func BytesToUint8(byt []byte) uint8 {
	return uint8(byt[0])
}

func Uint8ToBytes(i uint8, byt []byte) {
	byt[0] = byte(i)
}

`)

	for _, v := range tg.Uints[2:] {
		fmt.Println(tg.ConvUint(tg.Kind(v)))
	}
}
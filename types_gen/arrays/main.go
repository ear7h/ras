package main

import (
	"fmt"

	tg "github.com/ear7h/ras/types_gen"
)

func main() {
	fmt.Println("//this file is generated")
	fmt.Println("package ras\n")

	fmt.Println(`import (
	"encoding/binary"
)
`)

	for _, v := range tg.Uints {
		fmt.Println(tg.Array{
			Kind: v,
			Get: tg.UintArrGet(tg.BitSizeMap[v]),
			Set: tg.UintArrSet(tg.BitSizeMap[v]),
		})
	}

	for _, v := range tg.Ints {
		fmt.Println(tg.Array{
			Kind: v,
			Get: tg.IntArrGet(tg.BitSizeMap[v]),
			Set: tg.IntArrSet(tg.BitSizeMap[v]),
		})
	}

	//for _, v := range tg.Floats {
	//	fmt.Println(tg.Array{
	//		Kind: v,
	//		Get: tg.FloatGet(tg.BitSizeMap[v]),
	//		Set: tg.FloatSet(tg.BitSizeMap[v]),
	//	})
	//}
}

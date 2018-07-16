package main

import (
	"fmt"

	tg "github.com/ear7h/ras/types_gen"
)

func main() {
	fmt.Println("// this file is generated")
	fmt.Println("package ras\n")

	fmt.Println(`import (
	"encoding/binary"
	"math"
)
`)

	for _, v := range tg.Uints {
		fmt.Println(tg.Kind{
			Kind: v,
			Get: tg.UintGet(tg.BitSizeMap[v]),
			Set: tg.UintSet(tg.BitSizeMap[v]),
		})
	}

	for _, v := range tg.Ints {
		fmt.Println(tg.Kind{
			Kind: v,
			Get: tg.IntGet(tg.BitSizeMap[v]),
			Set: tg.IntSet(tg.BitSizeMap[v]),
		})
	}

	for _, v := range tg.Floats {
		fmt.Println(tg.Kind{
			Kind: v,
			Get: tg.FloatGet(tg.BitSizeMap[v]),
			Set: tg.FloatSet(tg.BitSizeMap[v]),
		})
	}

	//fmt.Println(tg.Kind{
	//	Kind: reflect.Bool,
	//	Get:tg.GetBool(),
	//	Set: tg.SetBool(),
	//})
}
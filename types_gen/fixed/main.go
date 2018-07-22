package main

import (
	"fmt"
	"reflect"

	tg "github.com/ear7h/ras/types_gen"
)

func main() {
	fmt.Print("// this file is generated\n")
	fmt.Print("package ras\n\n")

	fmt.Print(`//
// uints
//
`)
	for _, v := range tg.Uints[1:] {
		fmt.Println(tg.Kind(v))
	}

	fmt.Print(`//
// ints
//
`)
	for _, v := range tg.Ints[1:] {
		fmt.Println(tg.Kind(v))
	}

	fmt.Print(`//
// floats
//
`)
	for _, v := range tg.Floats {
		fmt.Println(tg.Kind(v))
	}

	fmt.Print(`//
// bool
//
`)
	fmt.Println(tg.Kind(reflect.Bool))
}
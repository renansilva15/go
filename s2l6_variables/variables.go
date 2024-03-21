package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v1 string      // Explicit.
	var v2 = "example" // Implicit.

	fmt.Println("Type of v1:", reflect.TypeOf(v1))
	fmt.Println("Type of v2:", reflect.TypeOf(v2))

	// Short variable declaration symbol.
	v3 := "example" // Implicit.

	fmt.Println("Type of v3:", reflect.TypeOf(v3))

	var (
		v4 string = "v4example"
		v5 string = "v5example"
	)

	v6, v7 := "v6example", "v7example"

	fmt.Println(v4, v5, v6, v7)

	const v8 = "v8example"

	fmt.Println(v8)
}

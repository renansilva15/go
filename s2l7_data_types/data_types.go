package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	// Integer types in Go:
	// int int8 int16 int32 int64

	// There are unsigned integers too:
	// uint uint8 uint16 uint32 uint64

	const biteType byte = 255 // Alias to uint8, cannot handle negative numbers.

	fmt.Println(reflect.TypeOf(biteType)) // uint8

	// When using float types, you need to choose between "float32" or "float64". However, if you use Go's type inference, it will default to "float", whose size depends on your system architecture.
	floatNumber := 0.0
	fmt.Println(reflect.TypeOf(floatNumber))

	// For all size of strings, use:
	const sentence string = "Hello World!"
	const letter string = "H"

	fmt.Println(sentence, letter)

	// In Go, you can only use single quotes with single characters. If you do that, Go converts it to a number which is the ASCII character code.
	const upperCaseAasciiOrder = 'A'

	fmt.Println(upperCaseAasciiOrder) // 65

	// Go has also an error type
	// To create a new error, you can use the native Go library "errors".
	var err error = errors.New("Some error")

	fmt.Println(err)

	// All of Go types has a "zero" value, you can see it if you define a variable with a type and do not assign a value to it.
	var number int
	var floatNumber2 float64
	var str string
	var boolean bool
	var err2 error

	fmt.Println(number)       // 0
	fmt.Println(floatNumber2) // 0
	fmt.Println(str)          // Empty string.
	fmt.Println(boolean)      // false
	fmt.Println(err2)         // <nil>
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("ARITHMETIC")
	const sum int = 20 + 10
	const subtraction int = 20 - 10
	const multiplication int = 20 * 10
	const division int = 20 / 10
	const remainder int = 20 % 10

	fmt.Println(sum, subtraction, division, multiplication, remainder)

	const n1 int = 20
	const n2 int32 = 10

	// You can't perform any operations with variables of different types.
	// sum2 := n1 + n2

	fmt.Println("ASSIGNMENT")
	var v1 = "example"
	v2 := "example"
	fmt.Println(v1, v2)

	fmt.Println("RELATIONAL OPERATORS")
	fmt.Println(1 > 1)
	fmt.Println(1 < 1)
	fmt.Println(1 >= 1)
	fmt.Println(1 <= 1)
	fmt.Println(1 == 1)
	fmt.Println(1 != 1)

	fmt.Println("LOGICAL OPERATORS")
	true, false := true, false
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

	fmt.Println("UNARY OPERATORS")
	number := 20
	fmt.Println(number)

	number++
	fmt.Println(number)

	number--
	fmt.Println(number)

	number += 10
	fmt.Println(number)

	number -= 10
	fmt.Println(number)

	number *= 10
	fmt.Println(number)

	number /= 10
	fmt.Println(number)

	number %= 10
	fmt.Println(number)

	// In Go there is no ternary operator.

	// In Go you can create an "if scoped" variable.
	var averageOfTheGrades = func(grade1, grade2 float64) (float64, error) {
		var err error
		var average float64

		if grade1 < 0 || grade1 > 10 || grade2 < 0 || grade2 > 10 {
			err = errors.New("Invalid grades.")
		} else {
			average = (grade1 + grade2) / 2
		}

		return average, err
	}

	if average, err := averageOfTheGrades(6, 8); err == nil {
		if average >= 7 {
			fmt.Println("Approved")

		} else {
			fmt.Println("Failed")
		}
	}
}

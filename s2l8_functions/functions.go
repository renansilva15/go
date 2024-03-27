package main

import "fmt"

func calculator(n1, n2 int) (int, int, int, int) {
	add := n1 + n2
	sub := n1 - n2
	mul := n1 * n2
	div := n1 / n2

	return add, sub, mul, div
}

func main() {
	fmt.Println(calculator(20, 10))

	_, sub, _, div := calculator(20, 10)

	fmt.Println(sub, div)

	var calculateMod = func(n1, n2 int) int {
		return n1 % n2
	}

	fmt.Println(calculateMod(20, 10))
}

package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(5, 1))

	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(2, 1))
}

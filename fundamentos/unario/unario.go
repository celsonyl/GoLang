package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix
	x++
	fmt.Println("X:", x)

	y--
	fmt.Println("Y:", y)

	//fmt.Println(x == y--) -> Comparação invalida
}

package main

import "fmt"

func main() {
	n := 1
	inc1(n) // passando por valor
	fmt.Println(n)

	inc2(&n)
	fmt.Println(n)
}

func inc1(n int) {
	n++
}

//Ponteiro é representado por *
func inc2(n *int) {
	*n++
}

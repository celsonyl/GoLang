package main

import "fmt"

func main() {
	x, y := 2, 3
	fmt.Println(x, y)
	fmt.Println(trocar(x, y))
}

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return
}

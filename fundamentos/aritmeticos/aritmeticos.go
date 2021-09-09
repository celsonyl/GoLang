package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma -> ", a+b)
	fmt.Println("Subtação -> ", a-b)
	fmt.Println("Divisão -> ", a/b)
	fmt.Println("Multiplicação -> ", a*b)
	fmt.Println("Módulo -> ", a%b)

	//bitwise
	fmt.Println("E ->", a&b)
	fmt.Println("Ou ->", a|b)
	fmt.Println("Xor ->", a^b)

	c := 3.0
	d := 2.0
	//outras op usando Math
	fmt.Println("Maior ->", math.Max(c, d)) //Float 64
	fmt.Println("Menor ->", math.Min(c, d)) //Float 64
	fmt.Println("Exponeciação ->", math.Pow(c, d))
}

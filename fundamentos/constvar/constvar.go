package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo (float64)

	//forma reduzida de criar var ( := )
	area := PI * math.Pow(raio, 2)
	fmt.Println("A area da Circunferencia é: ", area)

	const (
		a = 1
		b = 2
	)
	var (
		c = 4
		d = 5
	)
	fmt.Println(c, d)

	var e, f = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "Oloco"
	fmt.Println(g, h, i)
}

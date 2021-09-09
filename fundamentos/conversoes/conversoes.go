package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//Perigo
	fmt.Println("Teste " + string(rune(123)))

	//int para String
	fmt.Println("Teste " + strconv.Itoa(10))

	//String para int
	num, _ := strconv.Atoi("11") //retorna dois valores
	fmt.Println(num - 11)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}

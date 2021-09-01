package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é: ", reflect.TypeOf(3200))

	// Sem sinal(Só positivos). uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é: ", reflect.TypeOf(b))

	// Com sinal. int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é: ", i1)
	fmt.Println("O tipo do valor maximo é: ", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (uint32)
	fmt.Println("O rune é: ", reflect.TypeOf(i2))
	fmt.Println(i2)

	// Numeros reais float32 float64
	var x float32 = 49.99
	fmt.Println("O valor de x é: ", reflect.TypeOf(x))
	fmt.Println("O Tipo do literal é: ", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O valor do bo é: ", reflect.TypeOf(bo))

	// Tipo String
	s1 := "Hello world"
	fmt.Println("O tamanho da string é: ", len(s1))

	// String com varias linhas
	s2 := `Olá 
	testando
	varias
	linhas`
	fmt.Println("O tamanho é: ", len(s2))

	// Char não tem
	char := 'a'
	fmt.Println("O tipo do char é: ", reflect.TypeOf(char))
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, "Tipo ->", reflect.TypeOf(a1))
	fmt.Println(s1, "Tipo ->", reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array. Slide define um pedaço de um array
	s2 := a2[1:3] // Pegando o elemento '2 e 3' do array a2
	fmt.Println(s2)

	s3 := a2[:2] // Outro slice apontando para o mesmo array
	fmt.Println(s3)

	// Slice possui um tamanho e um ponteiro que aponta para um elemento de um array.
	s4 := s2[:1]
	fmt.Println(s4)
}

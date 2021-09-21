package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // Criando slice com 10 elementos mas a capacidade maxima internamente é 20
	fmt.Println(s, "Tamanho:", len(s), "Capacidade maxima:", cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, "Tamanho:", len(s))

	s = append(s, 1)  // Adicionando mais um elemento no slice
	fmt.Println(s, "Tamanho:", len(s), "Capacidade maxima:", cap(s)) // Quando estora a capcidade maxima ele dobra a capacidade ( Array Interno)
}

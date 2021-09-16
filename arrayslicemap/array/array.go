package main

import "fmt"

func main() {
	// homogênea (mesmo tipo de dado)
	var notas [3]float64
	notas[0], notas[1], notas[2] = 7.5, 5.3, 1.2
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Media %2.f", media)
}

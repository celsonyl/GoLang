package main

import "fmt"

func main() {
	imprimeResult(7)
	imprimeResult(5)
}

func imprimeResult(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com a nota:", nota)
	} else {
		fmt.Println("Reprovado com a nota:", nota)
	}
}

package main

import "fmt"

// Não tem operadore ternário em Go
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
	//return nota >= 6 ? "Aprovado" : "Reprovado"
}

func main() {
	fmt.Println(obterResultado(2))
}

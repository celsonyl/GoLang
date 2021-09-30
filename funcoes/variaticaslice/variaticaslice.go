package main

import "fmt"

func main() {
	aprovados := []string{"Celso", "Pedro", "Garrafa", "Bruno"}
	imprimirAprovador(aprovados...)
}

func imprimirAprovador(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d - %s ", i+1, aprovado)
	}
}

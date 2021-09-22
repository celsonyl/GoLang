package main

import "fmt"

func main() {
	funcSalarios := map[string]float64{
		"Celso": 1111.0,
		"Bruno": 1500.0,
		"Pedro": 5000.0,
	}

	funcSalarios["Garrafa"] = 1400.0
	delete(funcSalarios, "OK") // Tentar excluir um elemento que não existe. Não há problema
	for nome, salario := range funcSalarios {
		fmt.Printf("%s Salario:%.2f\n", nome, salario)
	}
}

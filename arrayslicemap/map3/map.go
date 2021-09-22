package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel":  1400.0,
			"Gabriela": 100.0,
			"Gustavo":  1400.0,
		},
		"C": {
			"Celso":  1400.0,
			"Carlos": 500.0,
		},
		"E": {
			"Eduardo Garrafa": 5000.0,
		},
	}

	for letra, nome := range funcPorLetra {
		fmt.Println(letra, nome)
	}
	delete(funcPorLetra, "C")

}

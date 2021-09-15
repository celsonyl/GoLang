package main

import "fmt"

func main() {
	fmt.Println(notaParaConceito(9))
	fmt.Println(notaParaConceito(7.3))
}

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 8 && nota < 9 {
		return "B"
	} else if nota >= 5 && nota < 8 {
		return "C"
	} else if nota >= 3 && nota < 5 {
		return "D"
	} else {
		return "E"
	}
}

package main

import "fmt"

func main() {
	fmt.Println(obterValorAprovado(5000))
	fmt.Println(obterValorAprovado(50))
}

func obterValorAprovado(n int) int {
	defer fmt.Println("fim!")

	if n >= 5000 {
		fmt.Println("Valor alto")
		return 5000
	} else {
		fmt.Println("Valor baixoo")
		return n
	}
}

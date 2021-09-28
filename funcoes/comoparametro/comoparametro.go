package main

import "fmt"

func main() {
	resultado := execute(multiplicacao, 3, 2)
	fmt.Println(resultado)
}

func multiplicacao(a, b int) int {
	return a * b
}

func execute(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

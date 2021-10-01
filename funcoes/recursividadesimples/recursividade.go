package main

import "fmt"

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)

	resultado2 := fatorial(-4) //->erro
	fmt.Println(resultado2)
}

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}

	//soma := 1
	//if n < 0 {
	//	return -1, fmt.Errorf("Numero invalido: %d", n)
	//} else if n == 0 {
	//	return 1, nil
	//} else {
	//	for i := 1; i <= n; n-- {
	//		aux := soma
	//		soma = aux * n
	//	}
	//	return soma, nil
	//}
}

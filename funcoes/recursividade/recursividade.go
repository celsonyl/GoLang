package main

import "fmt"

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Numero invalido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
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

package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { // não tem while em Go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	fmt.Println("\nMisturando")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par: ", i)
		} else {
			fmt.Println("Impar: ", i)
		}
	}

	for {
		// laço infinito
		fmt.Println("Para sempre")
		time.Sleep(time.Second)
	}
}

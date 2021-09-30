package main

import "fmt"

func main() {
	fmt.Printf("Media : %2.f", media(9, 12.0, 23.9))
}

func media(numeros ...float64) float64 { // ... -> podendo passar varios numeros
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

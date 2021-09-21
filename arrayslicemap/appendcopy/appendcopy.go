package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	var s1 []int

	//a1 = append(a1, 4, 5, 6) tem que passar um slice e nãp um array
	s1 = append(s1, 4, 5, 6)
	fmt.Println(a1, s1)

	s2 := make([]int, 2)
	copy(s2, s1) // Copia um slice pra outro
	fmt.Println(s2)
}

package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("Nova linha")

	x := 3.1415
	//fmt.Printf("O valor de X é:"+x)

	xs := fmt.Sprintln(x)
	fmt.Println("O valor de X é: " + xs)
	fmt.Println("O valor de X é: ", x)

	fmt.Printf("O valor de X é %.2f", x)

	a := 1
	b := 1.098223
	c := false
	d := "Olá"
	fmt.Printf("\n%d | %2.f | %t | %s", a, b, c, d)
}

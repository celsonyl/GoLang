package main

import "fmt"

func main() {
	fmt.Println("Main")
}

func init() { //-> Executa sem ser chamada e executa antes da main!
	fmt.Println("Inicio")
}

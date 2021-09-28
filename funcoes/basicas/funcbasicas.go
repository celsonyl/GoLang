package main

import "fmt"

func main() {
	f1()
	f2("Oi", "aa")
	fmt.Println(f3())
	f4("Celso", "Antonio")

	r51, r52 := f5()
	fmt.Printf("F5: %s\n", r51)
	fmt.Printf("F5: %s", r52)
}

func f1() {
	fmt.Println("F1: Hello world")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: Parametro 1 : %s | Parametro 2: %s\n", p1, p2)
}

func f3() string {
	return "F4: Função 3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("Função 4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

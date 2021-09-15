package main

import "fmt"

func main() {
	i := 1

	//Go não tem aritmética de ponteiros
	var p *int = nil

	p = &i          // pegando o endereço da variavel i
	fmt.Println(p)  //-> Printa o endereço que a variavel p aponta
	fmt.Println(*p) // -> Printa o valor de i

	fmt.Println(&i) // -> Printa o endereço de i
}

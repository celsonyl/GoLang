package main

import "fmt"

func main() {
	//var aprovados map[int]string -> maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456789] = "Celso"
	aprovados[122324441] = "Garrafa"
	aprovados[112312213] = "OK"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[123456789]) // Imprimir valor do map
	delete(aprovados, 123456789)      // -> Nome do map depois valor da key
	fmt.Println(aprovados[123456789], "Excluiu?")
}

package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nome string) {
	partes := strings.Split(nome, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{nome: "Pedro", sobrenome: "Silva"}
	p2 := pessoa{nome: "Pedro", sobrenome: "Silva"}

	fmt.Println(p1.getNomeCompleto())
	p2.setNomeCompleto("Celso Antonio")
	fmt.Println(p2.getNomeCompleto())
}

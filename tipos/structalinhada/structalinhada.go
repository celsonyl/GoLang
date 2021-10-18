package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valotTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{produtoID: 1, qtde: 2, preco: 10.14},
			{produtoID: 2, qtde: 1, preco: 9.14},
			{produtoID: 3, qtde: 5, preco: 14.14},
		},
	}

	fmt.Println(pedido)
	fmt.Println(pedido.valotTotal())
}

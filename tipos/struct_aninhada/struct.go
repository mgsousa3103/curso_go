package main

import "fmt"

type item struct {
	productId int
	qtde      int
	preco     float64
}

type pedido struct {
	userId int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}

	return total
}

func main() {
	pedido := pedido{
		userId: 1,
		itens: []item{
			{
				productId: 1,
				qtde:      5,
				preco:     10.64,
			},
			{
				productId: 2,
				qtde:      10,
				preco:     12.71,
			},
			{
				productId: 3,
				qtde:      7,
				preco:     8.55,
			},
		},
	}

	fmt.Printf("O valor total do pedido eh %.2f", pedido.valorTotal())
}

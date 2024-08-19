package main

import "fmt"

func pesquisaBinaria(item int, lista ...int) (int, string) {
	baixo := 0
	alto := len(lista) - 1

	for baixo <= alto {
		meio := (baixo + alto) / 2
		chute := lista[meio]

		if chute == item {
			return meio, "<== esta nesta posicao"
		}

		if chute > item {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}
	return -1, "nao existe este numero no array"
}

func main() {
	minha_lista := []int{1, 3, 5, 7, 9}

	fmt.Println(pesquisaBinaria(8, minha_lista...))
}

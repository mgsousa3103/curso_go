package main

import "fmt"

func ordenacaoPorSelecao(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		menor_indice := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[menor_indice] {
				menor_indice = j
			}
		}

		temp := arr[i]
		arr[i] = arr[menor_indice]
		arr[menor_indice] = temp
	}
	return arr
}

func main() {
	array := []int{5, 3, 6, 2, 10}

	fmt.Println(ordenacaoPorSelecao(array))
}

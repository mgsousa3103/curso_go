package main

import "fmt"

func main() {
	// var aprovados map[int]string -> mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456788] = "Maria"
	aprovados[981236378] = "Pedro"
	aprovados[743671930] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s, (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[123456788])
	delete(aprovados, 123456788)
	fmt.Println(aprovados)
}

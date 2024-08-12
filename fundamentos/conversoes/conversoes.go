package main

import (
	"fmt"
	conv "strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Teste " + conv.Itoa(123))

	// string para int
	num, _ := conv.Atoi("123")
	fmt.Println(num - 122)

	// conversao para booleal
	b, _ := conv.ParseBool("true") // 0 [false] e 1 [true]
	if b {
		fmt.Println("Verdadeiro")
	}
}

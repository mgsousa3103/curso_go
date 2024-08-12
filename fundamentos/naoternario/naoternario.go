package main

import "fmt"

// nao existe o operador ternario
func obterResultado(nota float64) string {
	// return nota >= 6 ? "Aprovado" : "Reprovado"

	if nota >= 6 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(5.4))
}

package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose Joao":      1153.55,
		"Gabriela Silva": 2553.55,
		"Pedro Junior":   4784.95,
	}

	funcsESalarios["Rafael Filho"] = 1277.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}

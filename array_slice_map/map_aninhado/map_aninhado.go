package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 1543.53,
			"Gustavo Lisboa": 7653.98,
		},
		"J": {
			"Jose Joao": 8376.53,
		},
		"P": {
			"Pedro Junior": 4231.76,
		},
	}

	delete(funcsPorLetra, "P")

	for _, funcs := range funcsPorLetra {
		for fs, sal := range funcs {
			fmt.Printf("%s, (Salario: %f)\n", fs, sal)
		}
	}
}

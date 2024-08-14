package main

import "fmt"

func notaParaConceito(nota float64) string {

	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota <= 9:
		return "B"
	case nota >= 5 && nota <= 8:
		return "C"
	case nota >= 3 && nota <= 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.6))
	fmt.Println(notaParaConceito(8.4))
	fmt.Println(notaParaConceito(7.8))
	fmt.Println(notaParaConceito(4.0))
	fmt.Println(notaParaConceito(2.7))
}

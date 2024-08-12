package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtracao =>", a-b)
	fmt.Println("Divisao =>", a/b)
	fmt.Println("Multiplicacao =>", a*b)
	fmt.Println("Modulo =>", a%b)

	// bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 = 10
	fmt.Println("OU =>", a|b)  // 11 | 10 = 11
	fmt.Println("XOR =>", a^b) // 11 ^ 10 = 01

	// operacoes utiliando a classe math
	c := 3.0
	d := 2.0

	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciacao =>", math.Pow(c, d))
}

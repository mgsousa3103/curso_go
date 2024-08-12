package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma variavel
	area := PI * m.Pow(raio, 2)
	fmt.Println("A area da circunferencia eh", area)

	// Outras formas de declarar const e var
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "Opa!"
	fmt.Println(g, h, i)
}

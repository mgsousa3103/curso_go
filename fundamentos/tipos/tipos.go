package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro eh", reflect.TypeOf(32000))

	// sem sinal (so positivos)... uint8 uint15 uint32 uint64
	var b byte = 255
	fmt.Println("O byte eh", reflect.TypeOf(b))

	// com sinal... uint8 uint16 uint32 uint64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int eh", i1)
	fmt.Println("O tipo de i1 eh", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela unicode (int32)
	fmt.Println("O rune eh", reflect.TypeOf(i2))
	fmt.Println(i2) // numero da tabela unicode

	// numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x eh", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 48.99 eh", reflect.TypeOf(49.99)) //float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo eh", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Ola meu nome eh Matheus"
	fmt.Println(s1 + "!!")
	fmt.Println("O tamanho da string s1 eh", len(s1))

	// string com multiplas linhas
	s2 := `Ola
	meu
	nome
	eh
	Matheus
	`
	fmt.Println("O tamanho da string s2 eh", len(s2))
}

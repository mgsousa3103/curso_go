package main

import "fmt"

func main() {
	i := 1 // 8 bytes de memoria

	var p *int = nil

	p = &i // guarda o end de memoria que aponta para o mesmo local que a var i esta apontando
	*p++   // desreferenciando (pegando o valor)
	i++

	// Go nao tem aritmetica de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)

}

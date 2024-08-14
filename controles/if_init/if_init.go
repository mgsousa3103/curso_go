package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	nano := rand.NewSource(time.Now().UnixNano())
	random := rand.New(nano)

	return random.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { // tambem suportado no switch
		fmt.Println("Ganhou !!")
	} else {
		fmt.Println("Perdeu !!")
	}
}

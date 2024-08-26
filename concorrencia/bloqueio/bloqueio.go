package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operacao bloqueante
	fmt.Println("So depois que foi lido")
}

func main() {
	ch := make(chan int) // canal sem buffer

	go rotina(ch)

	fmt.Println(<-ch) // operacao bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-ch) // deadlock
	fmt.Println("Fim")
}

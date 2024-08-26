package main

import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base
}

func main() {
	ch := make(chan int)
	go doisTresQuatroVezes(2, ch)

	a, b := <-ch, <-ch // recebendo dados do canal
	fmt.Println(a, b)

	fmt.Println(<-ch)
}

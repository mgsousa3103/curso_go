package main

import (
	"fmt"
	"time"
)

func main() {
	// for como while
	cont := 1
	for cont <= 5 { // nao tem while em Go
		fmt.Print(cont)
		cont++
	}

	fmt.Print("\n")

	// for como for
	for i := 0; i <= 8; i += 2 {
		fmt.Printf("%d", i)
	}

	// Misturando controles if, else, for
	fmt.Println("\nMisturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	// laco infinito
	for {
		fmt.Println("Para sempre..")
		time.Sleep(time.Second)
	}
}

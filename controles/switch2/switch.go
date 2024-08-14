package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()

	switch { // switch true
	case time.Hour() < 12:
		fmt.Println("Bom dia")
	case time.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}
}

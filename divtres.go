package main

import (
	"fmt"
)

func divtres(limite int) {
	for i := 1; i <= limite; i++ {
		if i%3 == 0 {
			fmt.Printf("O valor %v é divisível por 3\n", i)

		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	divtres(100)
}

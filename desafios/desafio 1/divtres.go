package main

import (
	"fmt"
)

// Cria função que detecta números divisíveis por 3
func divtres(limite int) {
	// Para cada i que começa em 1 e tem o limite definido pelo usuário...
	for i := 1; i <= limite; i++ {
		// Se i for divisível por 3...
		if i%3 == 0 {
			// Imprima a mensagem
			fmt.Printf("O valor %v é divisível por 3\n", i)

		} else {
			fmt.Println(i)
		}
	}
}

// Chama a função
func main() {
	divtres(100)
}

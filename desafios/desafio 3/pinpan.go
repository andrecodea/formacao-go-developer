package main

import (
	"fmt"
)

// Cria função no estilo FizzBuzz
func pinpan(limite int) {
	// Para cada i que começa em 1 e tem o limite definido pelo usuário
	for i := 1; i <= limite; i++ {
		// Se i for divisível por 3, imprima "Pin"
		if i%3 == 0 {
			fmt.Println("Pin")
		// Se i for divisível por 5, imprima "Pan"
		} else if i%5 == 0 {
			fmt.Println("Pan")
		// Se não, imprima i
		} else {
			fmt.Println(i)
		}
	}
}

// Chama a função
func main() {
	pinpan(100)
}

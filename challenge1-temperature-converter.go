// Construindo um conversor de temperaturas.
// - A conversão deve ser de Kelvin para Celsius.

// Iniciando o pacote principal do programa:
package main

// Importando fmt para formatação de entrada e saída:
import "fmt"

// Iniciando a função principal do programa:
func main() {
	// Criando a variável que receberá a temperatura em celsius:
	var kelvin float64

	// Criando o prompt para o usuário:
	fmt.Print("Digite a temperatura em Kelvin (K): ")

	// Recebendo a resposta do usuário:
	fmt.Scanf("%v", &kelvin)

	// Criando a variável de conversão:
	var celsius = kelvin - 273.15

	// Enviando a temperatura convertida:
	fmt.Printf("A temperatura equivalente em Celsius é: %v°C\n", celsius)

	// Criando features extras:

	/*Se a temperatura for 100°C, o usuário deverá ser informado
	que essa é a temperatura de ebulição da água, à titulo de curiosidade.*/
	if celsius == 100 {
		fmt.Println("Essa é a temperatura de ebulição da água!")
	}

	/*Se a temperatura for 0°C, o usuário deverá ser informado
	que essa é a temperatura de solidificação da água , à titulo de curiosidade.*/
	if celsius == 0 {
		fmt.Println("Essa é a temperatura de solidificação da água!")
	}
}

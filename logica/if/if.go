package condicionais

import "fmt"

// If é uma estrutura de controle de fluxo que permite executar um bloco de código se uma condição for verdadeira.
// Ele pode ser acompanhado por um bloco else para executar um código alternativo caso a condição seja falsa.
// As comparações em Go são feitas usando operadores de comparação, que são:
//  ==(igual), !=(diferente), >(maior), <(menor), >=(maior ou igual), <=(menor ou igual), ||(OU) e &&(E).
func If() {
	idade := 18
	if idade >= 18 {
		fmt.Println("Maior de idade")
	} else {
		fmt.Println("Menor de idade")
	}
}

// O else if é usado para verificar múltiplas condições em sequência, permitindo que você execute diferentes blocos de código
// com base em várias condições.
// Ele é útil quando você tem mais de duas opções para avaliar, proporcionando uma estrutura mais clara e organizada do que usar
// múltiplos if-else aninhados.
func IfElseIf() {
	nota := 85
	if nota >= 90 {
		fmt.Println("Aprovado com distinção")
	} else if nota >= 70 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}

// O if aninhado é uma estrutura de controle de fluxo onde um bloco if é colocado dentro de outro bloco if.
// Ele é usado para verificar condições adicionais dentro de um bloco if, permitindo uma lógica mais complexa e detalhada.
// No entanto, o uso excessivo de if aninhados pode tornar o código difícil de ler e manter, por isso é importante usá-los com moderação.
func IfAninhado() {
	idade := 20
	carteiraDeHabilitacao := true
	if idade >= 18 {
		if carteiraDeHabilitacao {
			fmt.Println("Pode dirigir")
		} else {
			fmt.Println("Não pode dirigir")
		}
	} else {
		fmt.Println("Menor de idade")
	}
}

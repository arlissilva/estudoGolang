package condicionais

import "fmt"

// Switch é uma estrutura de controle de fluxo que permite executar um bloco de código com base no valor de uma expressão.
// Ele é uma alternativa mais legível e eficiente ao uso de múltiplos if-else para comparar um valor com várias opções.
// O switch em Go é mais flexível do que em outras linguagens, pois permite comparar expressões complexas e não se limita
// apenas a valores constantes.

func Switch() {
	dia := "Segunda-feira"
	switch dia {
	case "Segunda-feira":
		fmt.Println("Hoje é segunda-feira")
	case "Terça-feira":
		fmt.Println("Hoje é terça-feira")
	case "Quarta-feira":
		fmt.Println("Hoje é quarta-feira")
	default:
		fmt.Println("Dia da semana desconhecido")
	}
}

// O switch sem expressão é uma forma de usar o switch para avaliar condições booleanas diretamente,
// sem a necessidade de comparar um valor específico.
func SwitchSemExpressao() {
	nota := 85
	switch {
	case nota >= 90:
		fmt.Println("Aprovado com distinção")
	case nota >= 70:
		fmt.Println("Aprovado")
	default:
		fmt.Println("Reprovado")
	}
}

// O fallthrough é uma palavra-chave em Go que permite que a execução continue para o próximo case no switch,
// mesmo que a condição do case atual seja verdadeira.
// Ele é usado para criar casos que compartilham um bloco de código comum ou para executar múltiplos casos consecutivos.
func SwitchComFallthrough() {
	dia := "Segunda-feira"
	switch dia {
	case "Segunda-feira":
		fmt.Println("Hoje é segunda-feira")
		fallthrough
	case "Terça-feira":
		fmt.Println("Hoje é terça-feira")
		fallthrough
	case "Quarta-feira":
		fmt.Println("Hoje é quarta-feira")
	default:
		fmt.Println("Dia da semana desconhecido")
	}
}

func SwitchComVariosCases() {
	dia := "Sábado"
	switch dia {
	case "Sábado", "Domingo":
		fmt.Println("Hoje é fim de semana")
	default:
		fmt.Println("Hoje é dia de semana")
	}
}

package condicionais

import "fmt"

func For() {
	// O loop for é a única estrutura de repetição em Go, e pode ser usado de várias maneiras para iterar sobre coleções,
	// executar um bloco de código várias vezes, ou criar loops infinitos.
	// O loop for é composto por três partes: a inicialização, a condição e a pós-incrementação.
	// A inicialização é executada apenas uma vez, no início do loop. A condição é avaliada antes de cada iteração,
	// e o loop continua enquanto a condição for verdadeira. A pós-incrementação é executada após cada iteração,
	// e geralmente é usada para atualizar a variável de controle do loop.
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func ForRange() {
	// O loop for range é uma variação do loop for que é usado para iterar sobre coleções, como arrays, slices, mapas e strings.
	// Ele fornece uma maneira conveniente de acessar os elementos de uma coleção sem a necessidade de usar um índice.
	// i é o índice do elemento atual, e v é o valor do elemento atual.
	// O loop for range continua até que todos os elementos da coleção tenham sido iterados.
	colecao := []string{"um", "dois", "tres"}
	for i, v := range colecao {
		fmt.Printf("Índice: %d, Valor: %s\n", i, v)
	}
}

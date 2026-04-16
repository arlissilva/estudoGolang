package funcoes

import "fmt"

//Para exportar uma função, ela deve começar com letra maiúscula.
func Funcoes(a, b int) (int, bool) {
	// Função que soma dois números
	soma := func(a int, b int) int {
		return a + b
	}

	// Chama a função e imprime o resultado
	resultado := soma(a, b)
	if resultado > 10 {
		fmt.Println("O resultado é maior que 10.")
		return resultado, true
	}
	fmt.Println("Resultado da soma:", resultado)
	return resultado, false
}

func Variadicas(a ...int) int {
	// Função que recebe um número variável de argumentos e retorna a soma
	soma := 0
	for _, v := range a {
		// O += é um operador de atribuição que adiciona o valor à variável existente.
		soma += v
	}
	return soma
}

func Closures() func(int) int {
	// Função que retorna outra função (closure ou função anônima)
	x := 10
	return func(y int) int {
		return x + y
	}
}

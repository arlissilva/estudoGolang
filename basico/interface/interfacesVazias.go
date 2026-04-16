package interfaces

import "fmt"

//interfaces vazias são interfaces que não possuem nenhum método definido.
// Elas são usadas para representar qualquer tipo de valor, pois todos os tipos implementam a interface vazia.
// A interface vazia é representada por "interface{}" em Go.
//Não é recomendado usar interfaces vazias em excesso, pois elas podem levar a código menos seguro e mais difícil de entender,
// já que não há restrições sobre os tipos de valores que podem ser atribuídos a uma variável do tipo interface vazia.

func InterfacesVazias() {
	var valor interface{}                                                 // Declarando uma variável do tipo interface vazia
	valor = "Olá, mundo!"                                                 // Atribuindo uma string à variável
	fmt.Println(valor)                                                    // Imprimindo o valor da variável
	fmt.Printf("O tipo da variavel é: %T e o valor é %v\n", valor, valor) // Imprimindo o tipo e o valor da variável
	valor = 42                                                            // Atribuindo um inteiro à variável
	fmt.Println(valor)                                                    // Imprimindo o valor da variável
	fmt.Printf("O tipo da variavel é: %T e o valor é %v\n", valor, valor) // Imprimindo o tipo e o valor da variável
}

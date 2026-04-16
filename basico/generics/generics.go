package generics

import "fmt"

type NumberType interface {
	~int | float64 // O operador ~ é usado para indicar que o tipo int pode ser usado como um int normal e não somente como um tipo definido.
}

type MyInt int

// Generics é um recurso que permite criar funções e tipos que podem trabalhar com diferentes tipos de dados,
// sem a necessidade de duplicar código para cada tipo específico. Em Go, os generics foram introduzidos na versão 1.18.
//O T é um tipo genérico que pode ser substituído por qualquer tipo específico quando a função é chamada.
//No caso o T pode ser um int ou um float64.
// A função Soma recebe um mapa onde as chaves são strings e os valores são do tipo genérico T.
//E no final retorna um tipo T.
func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// A função SomaInterface é semelhante à função Soma, mas em vez de usar um tipo genérico T diretamente, ela usa uma interface chamada NumberType.
// A interface NumberType é definida para aceitar tipos int e float64, o que permite que a função SomaInterface seja usada com ambos os tipos de dados.
func SomaInterface[T NumberType](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaGenerica() {
	m := map[string]int{
		"um":   1,
		"dois": 2,
		"tres": 3,
	}
	fmt.Println(Soma(m))
}

func SomaGenericaComInterface() {
	m := map[string]MyInt{
		"um":   1,
		"dois": 2,
		"tres": 3,
	}
	fmt.Println(SomaInterface(m))
}

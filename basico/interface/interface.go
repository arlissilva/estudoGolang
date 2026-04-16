package interfaces

import "fmt"

type Endereco struct {
	Logradouro string
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome   string
	Idade  int
	Ativo  bool
	Adress Endereco
}

// Interfaces são tipos de dados abstratos que definem um conjunto de métodos que um tipo
// deve implementar para ser considerado como uma implementação da interface.
type Pessoa interface {
	Desativar()
}

// Implementando a interface Pessoa para a struct Cliente
func (c *Cliente) Desativar() {
	c.Ativo = false
}

// A função Desativacao recebe um parâmetro do tipo Pessoa, que é uma interface,
// e chama o método Desativar da interface.
// Isso permite que qualquer tipo que implemente a interface Pessoa possa ser desativado usando essa função.
func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func Structs() {

	c := Cliente{Nome: "Alice", Idade: 30, Adress: Endereco{Logradouro: "Rua A", Cidade: "Cidade B", Estado: "Estado C"}}
	fmt.Println("Cliente:", c)
	c.Idade = 31 // Modificando o campo Idade da struct Cliente
	fmt.Println("Cliente atualizado:", c)
	c.Adress.Cidade = "Cidade D" // Modificando o campo Cidade da struct Cliente
	fmt.Println("Cliente atualizado:", c)
	// Desativando o cliente
	c.Desativar()
	fmt.Println("Cliente após desativação:", c)
	Desativacao(&c) // Desativando o cliente usando a função Desativacao
	fmt.Println("Cliente após desativação usando função:", c)
}

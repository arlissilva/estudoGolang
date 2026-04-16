package ponteirosandstructs

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
}

func NewCliente(nome string, idade int) *Cliente { // Função construtora para criar um novo cliente. Retorna um ponteiro para a struct Cliente.
	return &Cliente{Nome: nome, Idade: idade} // O operador & é usado para obter o endereço de memória da struct, criando um ponteiro para ela.
}

func (c *Cliente) andou() { // Ao colocar o ponteiro *, a função tem acesso ao valor original da struct, permitindo modificar os campos diretamente.
	// se fosse sem o ponteiro, a função receberia uma cópia da struct, e as modificações não afetariam o valor original.
	c.Nome = "Maria"
	fmt.Printf("%s andou\n", c.Nome)
}

func PonteirosAndStructs() {
	cliente := NewCliente("João", 30) //cliente é uma variável do tipo Cliente, que é uma struct.
	// Ela é inicializada com os valores "João" para o campo Nome e 30 para o campo Idade.

	cliente.andou()
	fmt.Println(cliente)
}

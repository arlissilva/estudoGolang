package structs

import "fmt"

type Endereco struct {
	Logradouro string
	Cidade     string
	Estado     string
}

//Adress é um campo do tipo Endereco dentro da struct Cliente
// permitindo que cada cliente tenha um endereço associado.
// Composição é um princípio de design onde uma struct contém outra struct
type Cliente struct {
	Nome   string
	Idade  int
	Adress Endereco
}

//metodos em structs são funções associadas a um tipo específico,
// permitindo que você defina comportamentos para os tipos personalizados.
func (c Cliente) Saudacao() string {
	return fmt.Sprintf("Olá, meu nome é %s e tenho %d anos.", c.Nome, c.Idade)
}

func Structs() {
	//Structs são tipos de dados compostos que agrupam campos relacionados.
	//Definindo uma struct para representar um ponto no espaço 2D

	c := Cliente{Nome: "Alice", Idade: 30, Adress: Endereco{Logradouro: "Rua A", Cidade: "Cidade B", Estado: "Estado C"}}
	fmt.Println("Cliente:", c)
	c.Idade = 31 // Modificando o campo Idade da struct Cliente
	fmt.Println("Cliente atualizado:", c)
	c.Adress.Cidade = "Cidade D" // Modificando o campo Cidade da struct Cliente
	fmt.Println("Cliente atualizado:", c)
	fmt.Println(c.Saudacao()) // Chamando o método Saudacao da struct Cliente
}

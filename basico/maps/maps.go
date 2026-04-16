package maps

import "fmt"

func Maps() {
	// Cria um mapa de inteiros
	salarios := make(map[string]int)

	// Adiciona alguns valores ao mapa
	salarios["Alice"] = 50000
	salarios["Bob"] = 60000
	salarios["Charlie"] = 55000

	// Acessa um valor no mapa
	fmt.Printf("Salário de Alice: %d\n", salarios["Alice"])

	// Remove um valor do mapa
	delete(salarios, "Bob")

	// Verifica se um valor existe no mapa
	if salario, ok := salarios["Bob"]; ok {
		fmt.Printf("Salário de Bob: %d\n", salario)
	} else {
		fmt.Printf("Salário de Bob não encontrado.\n")
	}
}

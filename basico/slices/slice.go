package main

import "fmt"

func main() {
	// Cria um slice de inteiros com 5 elementos
	slice := []int{1, 2, 3, 4, 5}
	slicezerado := []int{} // Slice vazio, sem elementos
	fmt.Println(slice)
	fmt.Println(len(slice)) // Imprime o comprimento do slice

	// Adiciona um elemento ao slice
	slice = append(slice, 6)
	// Imprime o slice, seu comprimento e capacidade
	// A capacidade é a quantidade de espaço alocado para o slice, que pode ser maior do que o comprimento atual.
	// Quando o comprimento atinge a capacidade, o Go dobra o espaço para o slice.
	fmt.Printf("Elementos do slice: %v\nTamanho: %d\nCapacidade: %d\n", slice, len(slice), cap(slice))

	// Remove o elemento no índice 2
	slice = append(slice[:2], slice[3:]...)
	fmt.Printf("Elementos do slice: %v\nTamanho: %d\nCapacidade: %d\n", slice, len(slice), cap(slice))

	slicezerado = append(slicezerado, 10, 20, 30) // Adiciona elementos ao slice vazio
	fmt.Printf("Elementos do slice vazio: %v\nTamanho: %d\nCapacidade: %d\n", slicezerado, len(slicezerado), cap(slicezerado))
	slicezerado = append(slicezerado, 40, 50) // Adiciona mais elementos ao slice vazio
	fmt.Printf("Elementos do slice vazio após adição: %v\nTamanho: %d\nCapacidade: %d\n", slicezerado, len(slicezerado), cap(slicezerado))

}

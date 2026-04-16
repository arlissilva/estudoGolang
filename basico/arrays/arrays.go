package arrays

import "fmt"

func Array() {
	//Array tem tamanho fixo.
	//Cria um array de inteiros com 5 elementos
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	// len pega o tamanho do array, mas como o array começa do 0, subtraímos 1
	// para obter o índice do último elemento.
	fmt.Println(len(arr) - 1) //resultado do len é 5, com o menos -1, o resultado é 4
	//que é o índice do último elemento do array.

	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}

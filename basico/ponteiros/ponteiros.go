package ponteiros

import "fmt"

func Ponteiros() {
	// Declarando um ponteiro para inteiro
	var home int
	// Declarando um ponteiro para ponteiro de inteiro
	var home2 = &home

	home = 42     // Atribuindo um valor à variável home
	home2 = &home // Atribuindo o endereço de home a home2

	fmt.Println("Valor de home:", home)              // Valor de home: 42
	fmt.Println("Valor de home2: usando *:", *home2) //O simbolo * exibe o valor armazenado no endereço de memória apontado por home2
	fmt.Println("Endereço de home:", &home)          //O simbolo & exibe o endereço de memória da variável home
	fmt.Println("Endereço de home2:", home2)         // exibe o endereço de memória da variável home2 que é o mesmo de home
	fmt.Println("Valor de home2 usando *:", *home2)  // o simbolo * exibe o valor armazenado no endereço de memória apontado por home2
	fmt.Println("Valor de home2 usando &:", &home2)  // o simbolo & exibe o endereço de memória da variável home2 que não é o mesmo de home, pois home2
	// é um ponteiro para home, ou seja, armazena o endereço de memória de home, mas tem seu próprio endereço de memória.
}

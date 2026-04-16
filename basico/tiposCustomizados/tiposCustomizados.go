package tiposCustomizados

import "fmt"

// Define o tipo Custom como um tipo inteiro personalizado
type CustomType int

// Atribui um valor inteiro ao tipo Custom e imprime o resultado
func Custom() {
    var c CustomType
    c = 10
    fmt.Println(c)
}

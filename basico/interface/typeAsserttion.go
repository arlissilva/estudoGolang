package interfaces

//Type Assertation é uma forma de definir o tipo de uma interface vazia, ou seja,
// quando temos uma variável do tipo interface{} e queremos acessar seus métodos ou propriedades,
// precisamos fazer uma type assertion para informar ao compilador qual é o tipo real da variável.
func TypeAssertion() {
	var i interface{} = "hello"
	println(i.(string))   // i.(string) é a type assertion, que informa ao compilador que i é do tipo string
	result, ok := i.(int) // result é o valor da variável i, e ok é um booleano que indica se a type assertion foi bem sucedida
	if ok {
		println(result) // se a type assertion foi bem sucedida, printa o valor de result
	} else {
		println("type assertion failed") // se a type assertion falhou, printa uma mensagem de erro
	}
}

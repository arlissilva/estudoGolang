package jsonbasico

import (
	"encoding/json"
	"fmt"
	"os"
)

// As tags JSON (como `json:"Numero"`) são usadas para especificar o nome do campo no JSON resultante.
// Isso é útil para garantir que os campos sejam nomeados de maneira consistente no JSON,
// mesmo que os nomes dos campos na estrutura Go sejam diferentes.
// alguns exemplos de tags JSON incluem:
// `json:"field_name"`: Especifica o nome do campo no JSON. Se omitido, o nome do campo na estrutura Go será usado.
// `json:"field_name,omitempty"`: Especifica o nome do campo no JSON e omite o campo se ele tiver um valor zero (como 0 para números,
// "" para strings, nil para ponteiros, etc.).
// `json:"-"`: Indica que o campo deve ser ignorado durante a codificação e decodificação JSON.
// `json:"field_name,string"`: Especifica que o campo deve ser codificado como uma string, mesmo que seja um tipo numérico ou booleano na estrutura Go.
type Conta struct {
	Numero int `json:"Numero"`
	Saldo  int `json:"Saldo"`
}

// ****************************************************************************
// A diferença entre json.Marshal e json.NewEncoder é que json.Marshal guarda o resultado em uma variavel,
// enquanto json.NewEncoder escreve diretamente para um destino, como um arquivo ou a saída padrão.
// json.Marshal é útil quando você precisa manipular o JSON como uma string ou armazená-lo em uma variável,
// enquanto json.NewEncoder é mais eficiente para escrever JSON diretamente para um fluxo de saída, como um arquivo ou a resposta HTTP.
// ****************************************************************************

// JsonBasico é uma função que demonstra como converter uma estrutura Go em JSON usando a função json.Marshal.
// Ela cria uma instância da estrutura Conta, converte essa instância em JSON e imprime o resultado como uma string.
// Se ocorrer um erro durante a conversão, a função irá gerar um pânico.

// A função json.Marshal é usada para converter a estrutura Conta em um formato JSON.
// O resultado é um slice de bytes, que é convertido para string antes de ser impresso.

func JsonBasico() {
	conta := Conta{
		Numero: 12345,
		Saldo:  1000,
	}
	jsonData, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))

}

// JsonEncode é uma função que demonstra como usar o json.NewEncoder para codificar uma estrutura Go diretamente para um destino,
// como a saída padrão (os.Stdout).
// Ela cria uma instância da estrutura Conta e usa json.NewEncoder para codificar essa instância em JSON e enviá-la para a saída padrão.
// Se ocorrer um erro durante a codificação, a função irá gerar um pânico.
func JsonEncode() {
	conta := Conta{
		Numero: 12345,
		Saldo:  1000,
	}
	jsonData, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

}

// JsonDecode é uma função que demonstra como usar a função json.Unmarshal para decodificar um JSON em uma estrutura Go.
// Ela define um JSON como um slice de bytes, cria uma variável do tipo Conta e usa json.Unmarshal para preencher
// essa variável com os dados do JSON.
// Se ocorrer um erro durante a decodificação, a função irá gerar um pânico. Por fim, a função imprime a estrutura Conta preenchida.
func JsonDecode() {
	jsonData := []byte(`{"Numero":12345,"Saldo":1000}`)
	var conta Conta
	err := json.Unmarshal(jsonData, &conta)
	if err != nil {
		panic(err)
	}
	fmt.Println(conta)
}

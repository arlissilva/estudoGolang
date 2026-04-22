package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] { // 1: Itera sobre os argumentos passados na linha de comando, ignorando o primeiro (que é o nome do programa)
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro na requisição: %v\n", err) // O Fprintf é usado para formatar a string de erro e imprimir no stderr
		}
		defer req.Body.Close()           // Garante que o corpo da resposta seja fechado após a leitura
		res, err := io.ReadAll(req.Body) // Lê o corpo da resposta
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
		}

		var data ViaCep
		err = json.Unmarshal(res, &data) // Decodifica o JSON da resposta para a struct ViaCep
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao decodificar o JSON: %v\n", err)
		}

		file, err := os.Create("cep.json") // Cria um arquivo chamado "cep.json" para escrever a resposta
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo: %v\n", err)
		}
		defer file.Close() // Garante que o arquivo seja fechado após a escrita

		_, err = file.WriteString(fmt.Sprintf("CEP: %s", data)) // Escreve a resposta JSON no arquivo
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		}
	}

}

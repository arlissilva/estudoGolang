package arquivos

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// CriarArquivo cria um arquivo chamado "exemplo.txt" e exibe uma mensagem de sucesso ou erro.
// os.Create é usado para criar um arquivo. Se o arquivo já existir, ele será truncado (conteúdo apagado).
// O defer arquivo.Close() garante que o arquivo seja fechado corretamente após a operação, mesmo que ocorra um erro.
func CriarArquivo() {
	arquivo, err := os.Create("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer arquivo.Close()

	fmt.Println("Arquivo criado com sucesso:", arquivo.Name())
}

// EscreverArquivo cria um arquivo chamado "exemplo.txt" e escreve o conteúdo "Conteúdo de exemplo" nele.
// arquivo.WriteString é usado para escrever uma string no arquivo. Se ocorrer um erro durante a escrita, ele será tratado e exibido.
func EscreverArquivo() {
	arquivo, err := os.Create("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer arquivo.Close()

	tamanho, err := arquivo.WriteString("Conteúdo de exemplo")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Conteúdo escrito com sucesso no arquivo:", arquivo.Name(), "Tamanho:", tamanho)
}

// LerArquivo abre o arquivo "exemplo.txt" e lê seu conteúdo, exibindo-o no console.
// os.Open é usado para abrir um arquivo existente. Se o arquivo não existir, ele retornará um erro.
// io.ReadAll é usado para ler todo o conteúdo do arquivo. O conteúdo é convertido para string antes de ser exibido.
// pode usar tambem o os.ReadFile("exemplo.txt") para ler o arquivo inteiro de uma vez, sem precisar abrir e fechar o arquivo manualmente.
func LerArquivo() {
	arquivo, err := os.Open("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer arquivo.Close()

	conteudo, err := io.ReadAll(arquivo)
	if err != nil {
		fmt.Println("Erro ao ler arquivo:", err)
		return
	}

	fmt.Println("Conteúdo do arquivo:", string(conteudo))
}

// LerBufIo abre o arquivo "exemplo.txt" e lê seu conteúdo linha por linha usando bufio.Reader, exibindo cada linha no console.
// bufio.NewReader é usado para criar um leitor bufferizado a partir do arquivo.
// O método Read é usado para ler um bloco de bytes do arquivo, e o loop continua até que o final do arquivo (EOF) seja alcançado.
// O conteúdo lido é convertido para string antes de ser exibido.
// Este método é útil para ler arquivos grandes ou quando se deseja processar o conteúdo linha por linha, em vez de ler tudo de uma vez.
// pode usar tambem o bufio.Scanner para ler o arquivo linha por linha de forma mais simples, sem precisar lidar com buffers manualmente.
// make([]byte, 10) cria um buffer de 10 bytes para leitura, ou seja ele irá ler 10 bytes do arquivo a cada iteração do loop.
// Buffer é um espaço de memória temporário usado para armazenar dados durante a leitura ou escrita.
func LerBufIo() {
	arquivo, err := os.Open("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer arquivo.Close()

	reader := bufio.NewReader(arquivo)
	buffer := make([]byte, 10) // Buffer para leitura
	for {
		tamanho, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Erro ao ler arquivo:", err)
			return
		}
		fmt.Print(string(buffer[:tamanho])) //:tamanho é o tamanho real dos bytes lidos, pois o buffer pode ser maior do que os dados lidos.
	}
}

package apihttp

import "net/http"

// http.NewRequest é uma função que cria uma nova solicitação HTTP. Ela recebe o método HTTP (como "GET", "POST", etc.),
// a URL para a qual a solicitação será enviada e um corpo opcional (que pode ser nil para solicitações sem corpo).
// req.Header.Set é um método que permite definir um cabeçalho específico para a solicitação HTTP.
// Ele recebe o nome do cabeçalho e o valor que deve ser atribuído a ele.
// http.Client é uma estrutura que representa um cliente HTTP. Ele é usado para enviar solicitações HTTP e receber respostas.
// O método Do é usado para enviar a solicitação e obter a resposta do servidor.
// resp.Body.Close() é uma chamada para fechar o corpo da resposta HTTP após o uso.
// Isso é importante para liberar recursos e evitar vazamentos de memória.

func HttpBasico() {
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		panic(err)
	}
	// Adicionando cabeçalhos à solicitação
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	// Enviando a solicitação
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

package apihttp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// http.ResponseWriter é uma interface que permite escrever a resposta HTTP que será enviada de volta ao cliente.
// http.Request representa a solicitação HTTP recebida do cliente, contendo informações como o método HTTP, URL, cabeçalhos e corpo da solicitação.
// w.Header().Set("Content-Type", "application/json") define o cabeçalho da resposta HTTP para indicar que o conteúdo é do tipo JSON.
// w.WriteHeader(http.StatusOK) define o status da resposta HTTP como 200 OK, indicando que a solicitação foi bem-sucedida.
// fmt.Fprintln(w, `{"message": "Hello, World!"}`) escreve o corpo da resposta HTTP, que é um objeto JSON contendo a mensagem "Hello, World!".
// .Methods("GET", "POST") especifica que a função de manipulação deve ser chamada apenas para solicitações HTTP GET e POST.
// http.ListenAndServe(":8001", r) inicia o servidor HTTP na porta 8001 e usa o roteador mux para lidar com as solicitações recebidas.
func ApiHttp() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(
		w http.ResponseWriter,
		r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"message": "Hello, World!"}`)
	}).Methods("GET", "POST")
	http.ListenAndServe(":8001", r)
}

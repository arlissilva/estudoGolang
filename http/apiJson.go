package apihttp

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// ApiJson é uma função que configura um manipulador HTTP para responder a solicitações JSON.
// mux:= mux.NewRouter() cria um novo roteador usando a biblioteca Gorilla Mux, que é uma ferramenta popular para criar rotas em aplicativos Go.
// mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) { ... }).Methods("POST") define um manipulador para a rota "/user" que responde a solicitações HTTP POST. Dentro do manipulador, um objeto User é criado e preenchido com dados de exemplo. O cabeçalho da resposta é configurado para indicar que o conteúdo é JSON, o status da resposta é definido como 200 OK e o objeto User é codificado em JSON e enviado como resposta.
// http.ListenAndServe(":8001", mux) inicia o servidor HTTP na porta 8001 e usa o roteador mux para lidar com as solicitações recebidas.
func ApiJson() {
	mux := mux.NewRouter()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		user := User{
			Name:  "John Doe",
			Email: "john@example.com",
			Age:   30,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}).Methods("POST")

	if err := http.ListenAndServe(":8001", mux); err != nil {
		log.Fatal(err)
	}
}

package contextEstudos

import (
	"log"
	"net/http"
	"time"
)

func ContextServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//ctx recebe o context da requisição, ja que toda http.Request tem um método Context()
	// que retorna o context associado a essa requisição.
	ctx := r.Context()
	log.Println("Request Iniciada")
	defer log.Println("Request Finalizada")

	//Aqui o context é cancelado em casos onde o usuario executa um control C ou fecha a aba do navegador,
	// ou seja, quando o cliente cancela a requisição.
	//Quando isso acontece, o case <-ctx.Done() é executado, indicando que a requisição foi cancelada,
	// e a função retorna sem processar a requisição.
	select {
	case <-ctx.Done():
		log.Println("Request Cancelada")
		http.Error(w, "Request Cancelada", http.StatusRequestTimeout)
		return
	case <-time.After(5 * time.Second):
		log.Println("Request Processada")
		w.Write([]byte("Request Processada"))
	}

}

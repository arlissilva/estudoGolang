package pacoteHttp

import (
	"fmt"
	"io"
	"net/http"
)

func HttpPersonalizado() {
	//Cria um cliente HTTP personalizado usando a estrutura http.Client. Isso permite configurar opções
	// específicas para as solicitações HTTP, como tempo limite, transporte, etc.
	Client := &http.Client{}

	//Cria uma solicitação HTTP personalizada, nesse caso GET, porem pode ser POST, PUT, DELETE, etc.
	// O terceiro parâmetro é o corpo da solicitação, que pode ser nil para solicitações sem corpo.
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	//Envia a solicitação usando o Cliente e a solicitação personalizada.
	// O método Do é usado para enviar a solicitação e obter a resposta do servidor.
	// Ele retorna a resposta e um erro, se houver.
	resp, err := Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

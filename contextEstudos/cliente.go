package contextEstudos

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func ContextClient() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081", nil)
	if err != nil {
		panic(err)
	}

	//http.DefaultClient.Do(req) envia a requisição HTTP criada e retorna a resposta (res) ou um erro (err).
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	//io.Copy(os.Stdout, res.Body) lê os dados do corpo da resposta (res.Body) e os escreve no stdout (os.Stdout),
	// que é a saída padrão do programa.
	io.Copy(os.Stdout, res.Body)
}

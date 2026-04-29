package pacoteHttp

import (
	"context"
	"io"
	"net/http"
	"time"
)

func HttpContext() {
	//context.Background() cria um context vazio, sem regras ou valores associados.
	// Ele é usado como ponto de partida para criar contextos derivados.
	ctx := context.Background()

	//context.WithTimeout(ctx, time.Second) cria um novo context que é derivado do context original (ctx)
	// e tem um timeout de 1 segundo.
	// Isso significa que qualquer operação associada a esse novo context deve ser concluída
	// dentro de 1 segundo, caso contrário, o context será cancelado automaticamente.
	// O defer exit() garante que a função exit seja chamada para liberar os recursos associados
	// ao context quando a função HttpContext terminar sua execução.
	ctx, exit := context.WithTimeout(ctx, time.Second)
	defer exit()

	//http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com", nil) cria uma nova
	// requisição HTTP do tipo GET para a URL "https://www.google.com",
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	//http.DefaultClient.Do(req) envia a requisição HTTP criada e retorna a resposta (resp) ou um erro (err).
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}

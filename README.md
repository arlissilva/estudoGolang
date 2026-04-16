# Estudo 📚

Este repositório é somente para estudos básicos de Go: pequenos experimentos, exemplos e exercícios para aprender idiomatismos, ferramentas e padrões. Não é destinado a produção nem a distribuir código pronto — use apenas para referência e aprendizado.



## Comandos mais usados (rápido) ⚙️
- Ver versão e ambiente
    ```
    go version
    go env
    ```
- Inicializar e gerenciar módulos
    ```
    go mod init github.com/usuario/repositorio
    go mod tidy
    go list -m all
    ```
- Adicionar/atualizar dependências
    ```
    go get pacote@versão
    ```
- Build (local)
    ```
    go build ./...                      # build de todos os pacotes
    go build -o bin/app ./cmd/app       # executável único
    ```
- Executar
    ```
    go run ./cmd/app
    go run main.go
    ```
- Testes
    ```
    go test ./... -v
    go test ./pkg/exemplo -run TestNome
    go test -bench . -benchmem
    go test ./... -cover
    go test -coverprofile=coverage.out && go tool cover -html=coverage.out
    ```
- Formatação e verificação estática
    ```
    go fmt ./...
    go vet ./...
    golangci-lint run   # opcional, se instalado
    ```
- Instalar binário (local / módulo)
    ```
    go install ./cmd/app
    ```

## Validar a arquitetura do seu SO
```
go env GOOS GOARCH
```

## Builds para diferentes SOs 🏗️🖥️
- Windows (exe)
    ```
    set GOOS=windows
    set GOARCH=amd64
    go build -o bin/app.exe ./cmd/app
    ```
  Ou (PowerShell):
    ```
    $env:GOOS="windows"; $env:GOARCH="amd64"; go build -o bin/app.exe ./cmd/app
    ```
- Linux (amd64)
    ```
    GOOS=linux GOARCH=amd64 go build -o bin/app-linux-amd64 ./cmd/app
    ```
- Linux (arm64)
    ```
    GOOS=linux GOARCH=arm64 go build -o bin/app-linux-arm64 ./cmd/app
    ```
- macOS (darwin amd64 / arm64)
    ```
    GOOS=darwin GOARCH=amd64 go build -o bin/app-darwin-amd64 ./cmd/app
    GOOS=darwin GOARCH=arm64  go build -o bin/app-darwin-arm64  ./cmd/app
    ```
- Build estático (útil para container Linux)
    ```
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/app-linux-amd64-static ./cmd/app
    ```
- Usando Docker (ambiente limpo / cross-build)
    ```
    docker run --rm -v "$PWD":/src -w /src golang:1.20 \
      /bin/bash -c "GOOS=linux GOARCH=amd64 go build -o bin/app ./cmd/app"
    ```
- Multi-arch / release (sugestão)
    - Use ferramentas como goreleaser ou docker buildx para gerar artefatos multi-arch automatizados.

## Observações 📝
- Arquitetura sugerida: cmd/, pkg/, internal/, tests/.
- Substitua URLs, nomes de pacotes e licença conforme o projeto.
- Mantenha commits pequenos e rode testes/linters antes de abrir PRs.
- Indique claramente no README quando houver exemplos incompletos ou código experimental. ⚠️

(Arquivo destinado a aprendizado: exemplos podem ser incompletos; use para referência.)

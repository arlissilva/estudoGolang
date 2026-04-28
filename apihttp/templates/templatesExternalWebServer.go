package templates

import (
	"net/http"
	"text/template"
)

// Define a struct EstadoTemplate com os campos Name e Uf para representar um estado.
type EstadoTemplateWebServer struct {
	Name string
	Uf   string
}

// Define um tipo EstadoTemplates que é um slice de EstadoTemplate para armazenar múltiplos estados.
// Igual a um array de estados, mas com tamanho dinâmico.
type EstadoTemplatesServer []EstadoTemplateWebServer

//Em HandleFunc criamos uma função anonima que recebe o w (ResponseWriter ou resposta) e r (Request ou requisição)
// como parâmetros. Dentro dessa função, carregamos o template usando template.Must e template.New,
// especificando o nome do template e o caminho para o arquivo. Em seguida, executamos o template
// passando o ResponseWriter (w) e os dados (EstadoTemplatesServer) para renderizar a resposta HTTP.
// Se ocorrer um erro durante a execução do template, a função irá panic.

// O template.New deve ter o mesmo nome do arquivo template.html que é o arquivo externo, caso contrário,
// ocorrerá um erro e o template não será executado.
func TemplateMustWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("template.html").ParseFiles("apihttp/templates/internal/template.html"))
		err := tmpl.Execute(w, EstadoTemplatesServer{
			{"São Paulo", "SP"},
			{"Rio de Janeiro", "RJ"},
		})
		if err != nil {
			panic(err)
		}
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}

}

package templates

import (
	"os"
	"text/template"
)

// Define a struct EstadoTemplate com os campos Name e Uf para representar um estado.
type EstadoTemplate struct {
	Name string
	Uf   string
}

// Define um tipo EstadoTemplates que é um slice de EstadoTemplate para armazenar múltiplos estados.
// Igual a um array de estados, mas com tamanho dinâmico.
type EstadoTemplates []EstadoTemplate

// A função TemplateMustEstados utiliza template.Must para carregar um template a partir de um arquivo
// e executá-lo com uma lista de estados. O template é carregado a partir do arquivo "template.html"
// localizado no diretório "apihttp/templates/internal". Se ocorrer um erro durante o carregamento ou
// execução do template, a função irá panic.

// O template.New deve ter o mesmo nome do arquivo template.html que é o arquivo externo, caso contrário,
// ocorrerá um erro e o template não será executado.
func TemplateMustEstados() {
	tmpl := template.Must(template.New("template.html").ParseFiles("apihttp/templates/internal/template.html"))
	err := tmpl.Execute(os.Stdout, EstadoTemplates{
		{"São Paulo", "SP"},
		{"Rio de Janeiro", "RJ"},
	})
	if err != nil {
		panic(err)
	}
}

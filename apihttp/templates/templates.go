package templates

import (
	"os"
	"text/template"
)

type Estado struct {
	Name string
	Uf   string
}

func TemplateEstado() {
	estado := Estado{Name: "São Paulo", Uf: "SP"}
	// Criando um template
	tmpl := template.New("TemplateEstado")                             // TemplateEstado é o nome do template
	tmpl, err := tmpl.Parse("O estado {{.Name}} tem a sigla {{.Uf}}.") // O template usa {{.Name}} e {{.Uf}} para acessar os campos do struct Estado
	if err != nil {
		panic(err)
	}

	// Executando o template com os dados do estado
	err = tmpl.Execute(os.Stdout, estado)
	if err != nil {
		panic(err)
	}
}

package templates

import (
	"html/template"
	"os"
)

type EstadoAtual struct {
	Name string
	Uf   string
}

// A diferença entre template.Must e template.New é que template.Must é uma função de conveniência que chama
// template.New e verifica se ocorreu um erro. Se ocorrer um erro, template.Must irá panic,
// enquanto template.New retornará o erro para que você possa lidar com ele. Portanto, template.Must
// é útil quando você tem certeza de que o template é válido e deseja simplificar o código, enquanto
// template.New é mais flexível e permite lidar com erros de forma mais controlada.
func TemplateMustEstado() {
	estado := EstadoAtual{"São Paulo", "SP"}
	tmpl := template.Must(template.New("TemplateEstado").Parse("O estado {{.Name}} tem a sigla {{.Uf}}."))
	err := tmpl.Execute(os.Stdout, estado)
	if err != nil {
		panic(err)
	}
}

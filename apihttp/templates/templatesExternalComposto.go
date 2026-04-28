package templates

import (
	"net/http"
	"text/template"
)

// Define a struct EstadoTemplate com os campos Name e Uf para representar um estado.
type EstadoTemplateComposto struct {
	Name string
	Uf   string
}

// Define um tipo EstadoTemplates que é um slice de EstadoTemplate para armazenar múltiplos estados.
// Igual a um array de estados, mas com tamanho dinâmico.
type EstadoTemplatesComposto []EstadoTemplateComposto

//Em HandleFunc criamos uma função anonima que recebe o w (ResponseWriter ou resposta) e r (Request ou requisição)
// como parâmetros. Dentro dessa função, carregamos o template usando template.Must e template.New,
// especificando o nome do template e o caminho para o arquivo. Em seguida, executamos o template
// passando o ResponseWriter (w) e os dados (EstadoTemplatesComposto) para renderizar a resposta HTTP.
// Se ocorrer um erro durante a execução do template, a função irá panic.

// O template.New deve ter o mesmo nome do arquivo template.html que é o arquivo externo, caso contrário,
// ocorrerá um erro e o template não será executado.
func TemplateMustComposto() {
	//templates é um slice de strings que contém os caminhos para os arquivos de template que serão usados para renderizar a página.
	// Esses arquivos são "header.html", "estados.html" e "footer.html", localizados no diretório "apihttp/templates/internal".
	// O uso de um slice permite que você carregue múltiplos templates de forma organizada e flexível, facilitando a manutenção
	// e a reutilização dos componentes do template.
	templates := []string{
		"apihttp/templates/internal/header.html",
		"apihttp/templates/internal/estados.html",
		"apihttp/templates/internal/footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//templates... é a sintaxe para passar o slice de strings como argumentos individuais para a função ParseFiles.
		//Ou seja o ... é usado para "desempacotar" o slice, permitindo que cada elemento do slice seja tratado como um argumento separado.
		//Em template.New o nome deve ser do arquivo principal do template, que é o "estados.html",
		// pois ele é o arquivo que contém a estrutura principal do template e os outros arquivos (header.html e footer.html)
		// são incluídos dentro dele.
		tmpl := template.Must(template.New("estados.html").ParseFiles(templates...))
		err := tmpl.Execute(w, EstadoTemplatesComposto{
			{"São Paulo", "SP"},
			{"Rio de Janeiro", "RJ"},
			{"Minas Gerais", "MG"},
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

package templates

// O pacote html/template é uma biblioteca da linguagem Go que fornece funcionalidades para criar e renderizar
// templates HTML de forma segura. Ele é projetado para evitar vulnerabilidades de injeção de código,
// como Cross-Site Scripting (XSS), ao escapar automaticamente os dados inseridos nos templates.
// O pacote html/template é amplamente utilizado para gerar páginas web dinâmicas, onde os dados são inseridos em um formato HTML,
//  garantindo que o conteúdo seja tratado de maneira segura e adequada para exibição em navegadores.

//Sempre use html/template e não text/template para renderizar templates HTML, pois o html/template é
// projetado para lidar com as peculiaridades e vulnerabilidades específicas do HTML.
import (
	"html/template"
	"net/http"
	"strings"
)

type EstadoTemplateFuncoes struct {
	Name string
	Uf   string
}

type EstadoTemplatesFuncoes []EstadoTemplateFuncoes

func PassarParaMaiusculo(s string) string {
	return strings.ToUpper(s)
}

func TemplateMustFuncoes() {
	//templates é um slice de strings que contém os caminhos para os arquivos de template que serão usados para renderizar a página.
	// Esses arquivos são "header.html", "estados.html" e "footer.html", localizados no diretório "apihttp/templates/internal".
	// O uso de um slice permite que você carregue múltiplos templates de forma organizada e flexível, facilitando a manutenção
	// e a reutilização dos componentes do template.
	templates := []string{
		"apihttp/templates/internal/estados.html",
		"apihttp/templates/internal/header.html",
		"apihttp/templates/internal/footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//templates... é a sintaxe para passar o slice de strings como argumentos individuais para a função ParseFiles.
		//Ou seja o ... é usado para "desempacotar" o slice, permitindo que cada elemento do slice seja tratado como um argumento separado.
		//Em template.New o nome deve ser do arquivo principal do template, que é o "estados.html",
		// pois ele é o arquivo que contém a estrutura principal do template e os outros arquivos (header.html e footer.html)
		// são incluídos dentro dele.

		tmpl := template.Must(template.New("estados.html").Funcs(template.FuncMap{
			"Maiusculo": PassarParaMaiusculo,
		}).ParseFiles(templates...))

		err := tmpl.ExecuteTemplate(w, "estados.html", EstadoTemplatesFuncoes{
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

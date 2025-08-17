package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

func simpleTemplate() {
	tmpl := `{{.}}`
	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(os.Stdout, "Hello, World!")
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	Age  int
	Name string
}

func structTemplate() {
	tmpl := `
	{{if gt .Age 20}}
	{{.Name}} is order than 20
	{{else}}
	{{.Name}} is not older than 20
	{{end}}
	`
	t := template.Must(template.New("").Parse(tmpl))
	user := User{Age: 21, Name: "Bob"}
	err := t.Execute(os.Stdout, user)
	if err != nil {
		log.Fatal(err)
	}
}

func rangeTemplate() {
	tmpl := `{{range .}} <p>{{.}}</p>{{end}}`
	t := template.Must(template.New("").Parse(tmpl))
	values := []string{"Hello", "World"}
	err := t.Execute(os.Stdout, values)
	if err != nil {
		log.Fatal(err)
	}
}

type Employee struct {
	Name string
}

type Company struct {
	Employees []Employee
}

func withTemplate() {
	tmpl := `
	{{with index .Employees 0}}
	{{.Name}}
	{{end}}
	`
	t := template.Must(template.New("").Parse(tmpl))
	company := Company{
		Employees: []Employee{
			{Name: "Bob"},
			{Name: "MIke"},
		},
	}
	err := t.Execute(os.Stdout, company)
	if err != nil {
		log.Fatal(err)
	}
}

func templateHtml() {
	tmpl := `<div>{{.}}</div>`
	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(os.Stdout, template.HTML(`<b>HTML</b>`))
	if err != nil {
		log.Fatal(err)
	}
}

func templateJs() {
	tmpl := `<script>{{.}}</script>`
	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(
		os.Stdout,
		template.JS(`alert("<script>1</script>")`),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func funcmap() {
	t := template.New("").Funcs(template.FuncMap{
		"FormatDateTime": func(format string, d time.Time) string {
			if d.IsZero() {
				return ""
			}
			return d.Format(format)
		}})
	tmpl := `{{FormatDateTime "2006年01月02日" .}}`
	t = template.Must(t.Parse(tmpl))
	err := t.Execute(os.Stdout, time.Now())
	if err != nil {
		log.Fatal(err)
	}
}

func defineTemplate() {
	t := template.Must(template.New("").ParseGlob("template/*.tmpl"))
	err := t.ExecuteTemplate(os.Stdout, "index", "これは本文です。")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defineTemplate()
}

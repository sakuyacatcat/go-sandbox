package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

type PageData struct {
	Title   string
	Heading string
	Content string
}

func main() {
	// Define the template
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Heading}}</h1>
		<p>{{.Content}}</p>
	</body>
	</html>
	`
	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(os.Stdout, PageData{
		Title:   "My Page",
		Heading: "Hello, World!",
		Content: "This is a simple HTML page.",
	})
	if err != nil {
		log.Fatalf("execution failed: %v", err)
	}

	// Define the template with range and index
	tmpl2 := `
	{{range .}}
	<p>{{.}}</p>{{end}}

	<p class="extract-1">{{index . 1}}</p>
	`
	t2 := template.Must(template.New("").Parse(tmpl2))
	values := []string{"Hello", "World"}
	err2 := t2.Execute(os.Stdout, values)
	if err2 != nil {
		log.Fatalf("execution failed: %v", err)
	}

	// Define the template with if-else
	tmpl3 := `
	{{if gt .Age 20}}
	{{.Name}} is older than 20
	{{else}}
	{{.Name}} is younger than 20
	{{end}}
	`
	t3 := template.Must(template.New("").Parse(tmpl3))
	type User struct {
		Name string
		Age  int
	}
	user := User{Name: "Alice", Age: 30}
	err3 := t3.Execute(os.Stdout, user)
	if err3 != nil {
		log.Fatalf("execution failed: %v", err)
	}

	// Define the template without escaping
	tmpl4 := `<script>{{.}}</script>`
	t4 := template.Must(template.New("").Parse(tmpl4))
	err4 := t4.Execute(os.Stdout, template.JS(`alert("<script>1</script>")`))
	if err4 != nil {
		log.Fatalf("execution failed: %v", err)
	}

	// Define the template with custom function
	t5 := template.New("").Funcs(template.FuncMap{
		"FormatDateTime": func(format string, d time.Time) string {
			if d.IsZero() {
				return ""
			}
			return d.Format(format)
		},
	})
	tmpl5 := `{{FormatDateTime "2006年 01月 02日" .}}`
	t5 = template.Must(t5.Parse(tmpl5))
	err5 := t5.Execute(os.Stdout, time.Now())
	if err5 != nil {
		log.Fatalf("execution failed: %v", err)
	}
}

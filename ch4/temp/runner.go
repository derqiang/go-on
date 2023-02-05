package _templ

import (
	"go_try/ch4/json/github"
	"html/template"
	"log"
	"os"
	"strings"
	"time"
)

type TextTempl int64
type HTMLTempl int64
type EscapeTmpl int64

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
func (tt *TextTempl) Run() {

	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	result, err := github.SearchIssues(strings.Split("repo:golang/go is:open json decoder", " "))
	if err != nil {
		log.Fatal(err)
	}

	report := template.Must(template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func (ht HTMLTempl) Run() {

	result, err := github.SearchIssues(strings.Split("repo:golang/go 3133 10535", " "))
	if err != nil {
		log.Fatal(err)
	}
	//jsonStr, err := json.MarshalIndent(result, "", "	")
	//log.Println(string(jsonStr))

	var issueList = template.Must(template.New("issuelist").Parse(`
		<h1>{{.TotalCount}} issues</h1>
		<table>
		<tr style='text-align: left'>
		  <th>#</th>
		  <th>State</th>
		  <th>User</th>
		  <th>Title</th>
		</tr>
		{{range .Items}}
		<tr>
		  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		  <td>{{.State}}</td>
		  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
		</tr>
		{{end}}
		</table>
	`))

	if err = issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func (et EscapeTmpl) Run() {
	const escape = `A: {{.A}}, B: {{.B}}`
	escapeTmpl := template.Must(template.New("escape").Parse(escape))
	var data struct {
		A string
		B template.HTML
	}
	data.A = "<H2>Hello!</H2>"
	data.B = "<H2>Hello!</H2>"
	if err := escapeTmpl.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

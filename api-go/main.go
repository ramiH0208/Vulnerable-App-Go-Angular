package main

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
)

var messages []string

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/comment", commentHandler)

    fmt.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := `<html><body>
    <h1>Commentaires</h1>
    <ul>
    {{range .}}<li>{{.}}</li>{{end}}
    </ul>
    <form action='/comment' method='POST'>
    <input name='msg'><button type='submit'>Envoyer</button></form>
    </body></html>`
    t := template.Must(template.New("index").Parse(tmpl))
    t.Execute(w, messages)
}

func commentHandler(w http.ResponseWriter, r *http.Request) {
    msg := r.FormValue("msg")
    messages = append(messages, msg)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

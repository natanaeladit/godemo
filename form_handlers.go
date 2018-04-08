package main

import (
	"fmt"
	"net/http"
	"path"
	"runtime"
	"text/template"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("No caller information")
	}

	tmpl := template.Must(template.ParseFiles(path.Dir(filename) + "/assets/index.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	// TODO do something with details if http MethodPost

	tmpl.Execute(w, struct{ Success bool }{true})
}

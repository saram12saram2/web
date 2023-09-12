package cmd

import (
	"html/template"
	"net/http"
)

type ErrorStatus struct {
	Code    int
	Message string
}

func CheckerErrors(w http.ResponseWriter, status int) {
	tmpl, err := template.ParseFiles("error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	Msg := ErrorStatus{status, http.StatusText(status)}
	err = tmpl.Execute(w, Msg)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}
}

package cmd

import (
	"net/http"
	"text/template"
	"web/asciiart"
)

var (
	word   string
	banner string
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// w.Header().Set("Allow", http.MethodGet)
		CheckerErrors(w, 405)
		// http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		CheckerErrors(w, 404)
		// http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		CheckerErrors(w, 500)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		CheckerErrors(w, 500)
		return
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		CheckerErrors(w, 500)
		return
	}
	if r.Method != http.MethodPost {
		CheckerErrors(w, 405)
		return
	}
	word = r.FormValue("word")
	banner = r.FormValue("banner")
	if len(banner) == 0 {
		banner = "standard"
	}

	result, StatusCode := asciiart.AsciiArt(word, banner)
	if StatusCode == 400 {
		CheckerErrors(w, 400)
		return
	}
	if StatusCode == 500 {
		CheckerErrors(w, 500)
		return

	}
	err = tmpl.Execute(w, result)
	if err != nil {
		CheckerErrors(w, 500)
		return
	}
}

package web

import (
	"html/template"
	"net/http"
	"net/url"
)

type Result struct {
	Output string
}

func BaseHandler(w http.ResponseWriter, page string, data Result) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/"+page,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	var data Result
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		font := r.FormValue("font")
		if text == "" || font == "" {
			http.Error(w, "404 Not Found", http.StatusNotFound)
			return
		}
		fontfile, err := LoadBanner(font)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		output := Render(text, fontfile)

		value := url.Values{}
		value.Add("art", output)

		http.Redirect(w, r, "/ascii-art?"+value.Encode(), http.StatusSeeOther)
		return
	}
	BaseHandler(w, "home.html", data)
}

func ResultHandler(w http.ResponseWriter, r *http.Request) {
	var data Result
	output := r.URL.Query().Get("art")
	data.Output = output
	BaseHandler(w, "ascii.html", data)
}

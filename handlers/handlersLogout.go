package handlers

import (
	"html/template"
	"net/http"
)

func HandleLogout(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "home.html", nil)
	}
}

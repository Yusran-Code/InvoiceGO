package service

import (
	"dasar-go/auth"
	"dasar-go/config"
	"dasar-go/model"
	"dasar-go/repository"
	"html/template"
	"net/http"
)

func HandleSetup(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := auth.GetSession(r)
		email := session.Values["email"].(string)

		if r.Method == http.MethodGet {
			profile, _ := repository.LoadUserProfileByEmail(config.DB, email)
			tmpl.ExecuteTemplate(w, "setup.html", profile)
			return
		}

		r.ParseForm()
		profile := model.AppProfile{
			Email:           email,
			NamaPT:          r.FormValue("nama_pt"),
			NamaBank:        r.FormValue("nama_bank"),
			NoRekening:      r.FormValue("no_rekening"),
			PenanggungJawab: r.FormValue("penanggung_jawab"),
			Alamat:          r.FormValue("alamat"),
			Kabupaten:       r.FormValue("kabupaten"),
		}

		if err := repository.SaveUserProfile(config.DB, profile); err != nil {
			http.Error(w, "Gagal simpan profil", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}

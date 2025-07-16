package handlers

import (
	"invoice-go/auth"
	"invoice-go/config"
	"invoice-go/model"
	"invoice-go/service"
	"html/template"
	"net/http"
)

func HandleSetup(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := auth.GetSession(r)
		email := session.Values["email"].(string)

		switch r.Method {
		case http.MethodGet:
			profile, _ := service.LoadProfileByEmail(config.DB, email)
			tmpl.ExecuteTemplate(w, "setup.html", profile)

		case http.MethodPost:
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

			if err := service.UpdateProfile(config.DB, profile); err != nil {
				http.Error(w, "gagal simpan profil", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/index", http.StatusSeeOther)
		}
	}
}

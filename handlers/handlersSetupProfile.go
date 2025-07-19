package handlers

import (
	"html/template"
	"invoice-go/auth"
	"invoice-go/config"
	"invoice-go/model"
	"invoice-go/service"
	"net/http"
)

func HandleSetup(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := auth.GetSession(r)
		email := session.Values["email"].(string)

		switch r.Method {
		case http.MethodGet:
			profile, err := service.LoadProfileByEmail(config.DB, email)
			if err != nil {
				// Jika tidak ditemukan, isi struct kosong (Email tetap diisi agar tidak null di form)
				profile = &model.AppProfile{Email: email}
			}
			if err := tmpl.ExecuteTemplate(w, "setup.html", profile); err != nil {
				http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
			}

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
				http.Error(w, "Gagal simpan profil: "+err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/index", http.StatusSeeOther)
		}
	}
}

package service

import (
	"dasar-go/auth"
	"dasar-go/config"
	"dasar-go/model"
	"dasar-go/repository"
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
)

func HandleHome(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "home.html", nil)
	}
}

func HandleIndex(tmpl *template.Template, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := auth.GetSession(r)
		email := session.Values["email"].(string)

		_, err := repository.LoadUserProfileByEmail(config.DB, email)
		if err != nil {
			http.Redirect(w, r, "/setup", http.StatusSeeOther)
			return
		}
		tmpl.ExecuteTemplate(w, "index.html", nil)
	}
}

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
func HandleGenerate(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()

			qty, _ := strconv.ParseFloat(r.FormValue("quantity_kg"), 64)
			dpp, _ := strconv.ParseFloat(r.FormValue("dpp"), 64)

			const hargaSatuan = 560.0
			const pengali = 3.0
			const hargaPokok = 595795.0

			displayQty := qty * hargaSatuan * pengali
			pokok := qty * hargaPokok
			ppn := pokok * 0.12
			total := pokok + ppn

			data := model.InvoiceData{
				InvoiceNumber: r.FormValue("invoice_number"),
				InvoiceDate:   r.FormValue("invoice_date"),
				Periode:       r.FormValue("periode"),
				QuantityKG:    qty,
				DisplayQty:    displayQty,
				DPP:           dpp,
				Pokok:         pokok,
				PPN:           ppn,
				Total:         total,
			}

			err := tmpl.ExecuteTemplate(w, "invoice.html", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		tmpl.ExecuteTemplate(w, "generate.html", nil)
	}
}

package auth

import (
	"context"
	"encoding/json"
	"invoice-go/config"
	"invoice-go/repository"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var Store *sessions.CookieStore

func InitSession() {
	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		panic("SESSION_KEY belum di-set atau gagal terbaca dari .env")
	}

	Store = sessions.NewCookieStore([]byte(sessionKey))
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
}

// Konfigurasi OAuth Google
var googleOauthConfig *oauth2.Config

func InitOAuthConfig() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

}

// REGISTER AUTH ROUTES
// =============================
func RegisterAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", handleLogin)
	mux.HandleFunc("/callback", handleCallback)
	mux.HandleFunc("/logout", handleLogout)
}

// LOGIN
// =============================
func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL("state-random")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// CALLBACK DARI GOOGLE
// =============================
func handleCallback(w http.ResponseWriter, r *http.Request) {
	log.Println("🚨 MASUK CALLBACK")

	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	log.Println("STATE:", state)

	if state != "state-random" {

		http.Error(w, "State tidak cocok", http.StatusBadRequest)
		return
	}

	// Tukar code jadi token
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {

		http.Error(w, "Gagal tukar code", http.StatusInternalServerError)
		return
	}

	// Ambil data user dari Google
	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {

		http.Error(w, "Gagal ambil data user", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {

		http.Error(w, "Gagal decode userInfo", http.StatusInternalServerError)
		return
	}

	email := userInfo["email"].(string)
	log.Println("✅ EMAIL DITEMUKAN")

	session, err := Store.Get(r, "session")
	if err != nil {

	}
	session.Values["authenticated"] = true
	session.Values["email"] = email

	if err := session.Save(r, w); err != nil {

	} else {
		log.Println("✅ SESSION TERSIMPAN")
	}

	profile, err := repository.GetUserEmail(config.DB, email)
	log.Println("PROFIL: DITEMUKAN")
	log.Println("ERROR LOAD PROFIL:", err)

	if err != nil || profile.NamaPT == "" {
		log.Println("➡️ Redirect ke /setup")
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
		return
	}

	log.Println("➡️ Redirect ke /index")
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

// LOGOUT

func handleLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

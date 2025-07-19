package auth

import (
	"encoding/json"
	"invoice-go/config"
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

func RegisterAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", handleLogin)
	mux.HandleFunc("/callback", handleCallback)
	mux.HandleFunc("/logout", handleLogout)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	returnTo := r.URL.Query().Get("returnTo")
	if returnTo == "" {
		returnTo = "/index"
	}

	session, _ := Store.Get(r, "session")
	session.Values["returnTo"] = returnTo
	session.Save(r, w)

	url := googleOauthConfig.AuthCodeURL("state-random")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Kode tidak ditemukan di callback", http.StatusBadRequest)
		return
	}

	token, err := googleOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Gagal menukar kode dengan token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := googleOauthConfig.Client(r.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Gagal mendapatkan user info dari Google: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Gagal decode user info", http.StatusInternalServerError)
		return
	}

	email, ok := userInfo["email"].(string)
	if !ok || email == "" {
		http.Error(w, "Email tidak tersedia", http.StatusInternalServerError)
		return
	}

	// ✅ Simpan ke session
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, "Gagal ambil session", http.StatusInternalServerError)
		return
	}

	session.Values["authenticated"] = true
	session.Values["email"] = email
	if err := session.Save(r, w); err != nil {
		http.Error(w, "Gagal simpan session", http.StatusInternalServerError)
		return
	}

	// ✅ Cek apakah user sudah pernah isi profil
	var exists bool
	err = config.DB.QueryRow(`SELECT EXISTS (SELECT 1 FROM user_profile WHERE email = $1)`, email).Scan(&exists)
	if err != nil {
		http.Error(w, "Gagal cek profil user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ Redirect sesuai kondisi
	if exists {
		http.Redirect(w, r, "/invoice.html", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
	}
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

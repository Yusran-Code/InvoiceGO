package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// MIDDLEWARE AUTH
// =============================

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "session")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func GetSession(r *http.Request) (*sessions.Session, error) {
	return Store.Get(r, "session")
}

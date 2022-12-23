package auth

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	//TODO remove this when LDAP/AD Auth
	if username == "root" && password == "root" {
		// set session cookie
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: "authenticated",
		})
		// return success response
		fmt.Fprint(w, "Login successful!")
	} else {
		// return error response
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}

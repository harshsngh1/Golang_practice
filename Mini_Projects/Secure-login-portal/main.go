package main

import (
	"fmt"
	"net/http"
	"time"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

// key is the username and value is the Login struct
var users = map[string]Login{}

func main() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/protected", protected)
	http.ListenAndServe(":8080", nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid request method", er)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	if len(username) < 8 || len(password) < 8 {
		er := http.StatusNotAcceptable
		http.Error(w, "Username and password must be at least 8 characters long", er)
		return
	}

	if _, ok := users[username]; ok {
		er := http.StatusConflict
		http.Error(w, "Username already exists", er)
		return
	}

	hashedPassword, _ := hashPassword(password)
	users[username] = Login{
		HashedPassword: hashedPassword,
	}
	fmt.Fprintf(w, "User %s registered successfully", username)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid request method", er)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, ok := users[username]
	if !ok || !checkPasswordHash(password, user.HashedPassword) {
		er := http.StatusUnauthorized
		http.Error(w, "Invalid username or password", er)
		return
	}

	sessionToken := generateSessionToken(32)
	// CSRF token generation
	csrfToken := generateSessionToken(32)

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour), // Set cookie to expire in 24 hours
		HttpOnly: true,                           // Prevent JavaScript access to the cookie
	})
	// Set CSRF token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour), // Set cookie to expire in 24 hours
		HttpOnly: false,                          // Allow JavaScript access to the CSRF token
	})

	//store session token in user struct
	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	users[username] = user
	fmt.Fprintf(w, "User %s logged in successfully", username)
}

func protected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid request method", er)
		return
	}
	if err := Authorize(r); err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized", er)
		return
	}
	username := r.FormValue("username")
	fmt.Fprintf(w, "CSRF Validation successful!, Welcome %s!", username)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if err := Authorize(r); err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized", er)
		return
	}

	// Clear session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // Set cookie to expire in the past
		HttpOnly: true,                           // Prevent JavaScript access to the cookie
	})
	// Clear CSRF token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // Set cookie to expire in the past
		HttpOnly: false,                          // Allow JavaScript access to the CSRF token
	})

	// Clear session token and CSRF token from user struct
	user, ok := users[r.FormValue("username")]
	if ok {
		user.SessionToken = ""
		user.CSRFToken = ""
		users[r.FormValue("username")] = user
	}
	fmt.Fprintf(w, "User %s logged out successfully", r.FormValue("username"))
}

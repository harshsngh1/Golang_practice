package main

import (
	"errors"
	"net/http"
)

var AuthErr = errors.New("Unauthorized")

func Authorize(r *http.Request) error {
	username := r.FormValue("username")
	user, ok := users[username]
	if !ok {
		return AuthErr
	}
	//get session token from cookie
	// and compare it with the one stored in the user struct
	// if they are not the same, return an error
	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		return AuthErr
	}

	//Get CSRF token from the header and compare it with the one stored in the user struct
	// if they are not the same, return an error
	csrfToken := r.Header.Get("X-CSRF-Token") // CSRF token generation
	if csrfToken == "" || csrfToken != user.CSRFToken {
		return AuthErr
	}
	return nil
}

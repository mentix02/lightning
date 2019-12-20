package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexandrevicenzi/unchained/pbkdf2"
	"net/http"
)

func getAuthorBookmarkedArticlesPks(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	username, err := authorizedRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(DetailResponse{err.Error()})
	} else {
		_ = json.NewEncoder(w).Encode(articleIdsSortedByAuthorBookmarks(username).toSlice())
	}

}

func getAuthTokenFromUsernameAndPassword(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	switch r.Method {

	case "POST":

		// Get POST form data.
		_ = r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if len(username) != 0 {
			if len(password) != 0 {
				hashedPassword, err := getHashedPasswordFromUsername(username)
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					_ = json.NewEncoder(w).Encode(DetailResponse{err.Error()})
				} else {
					valid, _ := pbkdf2.NewPBKDF2SHA256Hasher().Verify(password, hashedPassword)
					if valid {
						token, _ := getTokenFromUsername(username)
						_ = json.NewEncoder(w).Encode(map[string]string{"token": token})
						return
					}
					w.WriteHeader(http.StatusUnauthorized)
					_ = json.NewEncoder(w).Encode(DetailResponse{"Invalid credentials."})
					return
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				_ = json.NewEncoder(w).Encode(DetailResponse{"Field 'password' not provided."})
				return
			}
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			_ = json.NewEncoder(w).Encode(DetailResponse{"Field 'username' not provided."})
			return
		}
		break

	default:
		detail := fmt.Sprintf("Method \"%s\" not allowed.", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(DetailResponse{detail})

	}

}

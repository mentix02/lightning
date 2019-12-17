package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexandrevicenzi/unchained"
)

func getAuthorBookmarkedArticlesPks(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	username, err := authorizedRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(DetailResponse{err.Error()})
	} else {
		_ = json.NewEncoder(w).Encode(ArticleIdsSortedByAuthorBookmarks(username).toSlice())
	}

}

func getAuthTokenFromUsernameAndPassword(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "POST" {
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
					valid, err := unchained.CheckPassword(password, hashedPassword)
					check(err)
					if valid {
						token, _ := getTokenFromUsername(username)
						_ = json.NewEncoder(w).Encode(map[string]string{"token": token})
					} else {
						w.WriteHeader(http.StatusUnauthorized)
						_ = json.NewEncoder(w).Encode(DetailResponse{"Invalid credentials."})
					}
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				_ = json.NewEncoder(w).Encode(DetailResponse{"Field 'password' not provided."})
			}
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			_ = json.NewEncoder(w).Encode(DetailResponse{"Field 'username' not provided."})
		}
	} else {
		detail := fmt.Sprintf("Method \"%s\" not allowed.", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(DetailResponse{detail})
	}
}

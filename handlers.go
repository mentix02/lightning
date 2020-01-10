package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexandrevicenzi/unchained/pbkdf2"
)

func getAuthorBookmarkedArticlesPksHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	username, err := authorizedRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(DetailResponse{err.Error()})
	} else {
		_ = json.NewEncoder(w).Encode(articleIdsSortedByAuthorBookmarks(username).toSlice())
	}

}

func getAuthTokenFromUsernameAndPasswordHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	switch r.Method {

	case "POST":

		// Get POST form data.
		_ = r.ParseForm()

		// Verify required POST data exists.
		err := formBodyContainsKeys(r, []string{"username", "password"})
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			_ = json.NewEncoder(w).Encode(DetailResponse{"Field '" + err.Error() + "' not provided."})
			return
		}

		// This code will only reach if both the username
		// and password fields were verified to exist on #L35.
		username := r.Form.Get("username")
		password := r.Form.Get("password")

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

	default:
		detail := fmt.Sprintf("Method \"%s\" not allowed.", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(DetailResponse{detail})

	}

}

func getRecentArticlesHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	_ = json.NewEncoder(w).Encode(DetailResponse{"Invalid credentials."})
}

// Backend code for the APIs that powers The Medialist.
package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/bookmark/list/", getAuthorBookmarkedArticlesPksHandler)
	router.HandleFunc("/authors/authenticate/", getAuthTokenFromUsernameAndPasswordHandler)

	log.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))

}

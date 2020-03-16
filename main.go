// Backend code for the APIs that powers The Medialist.
package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	openLogFile("dev.log")
	log.SetFlags(log.Ldate | log.Ltime)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/articles/recent/", getRecentArticlesHandler)
	router.HandleFunc("/articles/detail/{slug}/", getArticleDetailHandler)
	router.HandleFunc("/bookmark/list/", getAuthorBookmarkedArticlesPksHandler)
	router.HandleFunc("/authors/authenticate/", getAuthTokenFromUsernameAndPasswordHandler)

	log.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", requestLogger(router)))

}

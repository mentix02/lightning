package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Enables cross origin control access for a handler.
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// Authorization headers are in the format of -
// 		"Token {key}"
// The string should be split with a whitespace and the
// second part should be return. There are a few more checks
// can be implemented but this works just fine for almost
// all check cases.
func extractKeyFromToken(token string) (string, error) {
	splitString := strings.Split(token, " ")
	if len(splitString) == 2 {
		if splitString[0] == "Token" {
			return splitString[1], nil
		}
		return "", errors.New("Invalid token.")
	}
	return "", errors.New("Authentication credentials were not provided.")
}

// Helper method to make sure a handler is authenticated.
func authorizedRequest(r *http.Request) (string, error) {
	headerToken := r.Header.Get("Authorization")
	key, err := extractKeyFromToken(headerToken)
	if err != nil {
		return "", err
	}
	username, err := getAuthorUsernameFromKey(key)
	if err == nil {
		return username, nil
	}
	return "", err
}

// Receives a "map" like object (an HTTP POST request Form field)
// and a list of strings acting as values for the map object to
// determine whether they all exist in the map. If not, return an
// error with the name of the field, else, return nil
func formBodyContainsKeys(r *http.Request, keys []string) error {
	var value string
	for _, key := range keys {
		value = r.Form.Get(key)
		if value == "" {
			return errors.New(key)
		}
	}
	return nil
}

func openLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}

		log.SetOutput(lf)
	}
}

func requestLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

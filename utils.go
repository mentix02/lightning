package main

import (
	"errors"
	"log"
	"net/http"
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

func logRequest(r *http.Request) {
	log.Println(r.Method + " request at " + r.URL.Path)
}

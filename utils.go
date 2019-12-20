package main

import (
	"errors"
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

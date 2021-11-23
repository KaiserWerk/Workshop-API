package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func authV1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		apiKey := r.Header.Get("X-Api-Key")
		if apiKey == "" {
			http.Error(w, "Authentication failed (API Key missing)", http.StatusUnauthorized)
			return
		}

		if !authenticateV1(apiKey) {
			http.Error(w, "Authentication failed (API Key invalid)", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func authV2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Authentication failed (header missing)", http.StatusBadRequest)
			return
		}

		parts := strings.Split(header, " ")
		if len(parts) != 2 {
			http.Error(w, "Authentication failed (malformed Authorization header)", http.StatusBadRequest)
			return
		}

		userPw, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			http.Error(w, "Authentication failed (base64 decode failed)", http.StatusInternalServerError)
			return
		}

		parts = strings.Split(string(userPw), ":")
		if len(parts) != 2 {
			http.Error(w, "Authentication failed (malformed value)", http.StatusBadRequest)
			return
		}

		if !authenticateV2(parts[1]) {
			http.Error(w, "Authentication failed (API Key invalid)", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

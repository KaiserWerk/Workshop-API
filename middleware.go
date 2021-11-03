package main

import "net/http"

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

		next.ServeHTTP(w, r)
	})
}

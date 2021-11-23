package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type AccessToken struct {
	AccessToken string    `json:"access_token"`
	Validity    time.Time `json:"validity"`
}

func authenticateV2Handler(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("X-Api-Token")
	if h == "" {
		http.Error(w, "missing X-Api-Token header", http.StatusUnauthorized)
		return
	}

	tempToken, err := loginV2(h)
	if err != nil {
		http.Error(w, "could not log in", http.StatusUnauthorized)
		return
	}

	at := AccessToken{
		AccessToken: tempToken,
		Validity:    time.Now().Add(time.Hour),
	}

	json, err := json.MarshalIndent(at, "", "  ")
	if err != nil {
		http.Error(w, "could not marshal json", http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

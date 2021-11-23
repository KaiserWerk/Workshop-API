package main

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

type user struct {
	name           string
	apiKey         string
	tempKey        string
	tempKeyExpires time.Time
}

var (
	userMut sync.Mutex
	users   = map[int]user{
		0: user{
			name:   "Robin K.",
			apiKey: "T1noKRpZOOKFMmVYifoczdL4sqgPAOq3vGtTr6WF",
		},
		1: user{
			name:   "Tobias M.",
			apiKey: "dcq817XiNUfUWz74E6fe9kINmNninZvEXoZIYHup",
		},
		2: user{
			name:   "Leo E.",
			apiKey: "RP1JYxo67LNJDNJeqTBoMIoxQJhVIm8f9tszeru2",
		},
	}
)

func authenticateV1(key string) bool {
	userMut.Lock()
	defer userMut.Unlock()

	for _, v := range users {
		if v.apiKey == key {
			return true
		}
	}

	return false
}

func generateTempKey() (string, error) {
	b := make([]byte, 40)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func loginV2(key string) (string, error) {
	userMut.Lock()
	defer userMut.Unlock()

	tempKey, err := generateTempKey()
	if err != nil {
		return "", nil
	}
	//user := user{}
	for k, v := range users {
		if v.apiKey == key {
			v.tempKey = tempKey
			v.tempKeyExpires = time.Now().Add(time.Hour)

			users[k] = v
		}
	}

	return tempKey, nil
}

func authenticateV2(key string) bool {
	userMut.Lock()
	defer userMut.Unlock()

	for _, v := range users {
		if v.tempKey == key {
			if !v.tempKeyExpires.After(time.Now()) {
				return false
			}
			return true
		}
	}

	return false
}

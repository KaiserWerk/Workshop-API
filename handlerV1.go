package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func productGetAllHandler(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(getAllProducts())
	if err != nil {
		http.Error(w, "could not send product list", http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func productGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	u, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, "could not determine Id", http.StatusInternalServerError)
		return
	}
	prod, err := getProduct(uint32(u))
	if err != nil {
		http.Error(w, "could not get product", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(prod)
	if err != nil {
		http.Error(w, "could not marshal JSON", http.StatusInternalServerError)
		return
	}
}

func productAddHandler(w http.ResponseWriter, r *http.Request) {

}

package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func productGetAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	json, err := json.MarshalIndent(getAllProducts(), "", "  ")
	if err != nil {
		http.Error(w, "could not send product list", http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func productGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
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

	json, err := json.MarshalIndent(prod, "", "  ")
	if err != nil {
		http.Error(w, "could not marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func productAddHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "could not unmarshal JSON", http.StatusBadRequest)
		return
	}

	p = addProduct(p)

	json, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		http.Error(w, "could not marshal JSON", http.StatusBadRequest)
		return
	}

	w.Write(json)
}

func productEditHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "could not unmarshal JSON", http.StatusBadRequest)
		return
	}

	editProduct(p)
}

func productRemoveHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	id := vars["id"]
	u, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, "could not determine Id", http.StatusInternalServerError)
		return
	}

	err = removeProduct(uint32(u))
	if err != nil {
		http.Error(w, "could not find product", http.StatusNotFound)
		return
	}
}

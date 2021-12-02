package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func productReviewGetAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header.Set("Content-Type", "application/json")
	defer r.Body.Close()

	vars := mux.Vars(r)
	id := vars["productid"]
	u, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, "could not determine Id", http.StatusInternalServerError)
		return
	}

	json, err := json.MarshalIndent(GetAllReviews(uint32(u)), "", "  ")
	if err != nil {
		http.Error(w, "could not send product list", http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func productReviewGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header.Set("Content-Type", "application/json")
	defer r.Body.Close()
	vars := mux.Vars(r)

	prodId, err := strconv.ParseUint(vars["productid"], 10, 32)
	if err != nil {
		http.Error(w, "could not determine Id", http.StatusBadRequest)
		return
	}

	revId, err := strconv.ParseUint(vars["reviewid"], 10, 32)
	if err != nil {
		http.Error(w, "could not determine Id", http.StatusBadRequest)
		return
	}

	review, err := GetReview(uint32(revId))
	if err != nil {
		http.Error(w, "could not find review by Id", http.StatusNotFound)
		return
	}

	if review.ProductId != uint32(prodId) {
		http.Error(w, "could not find review by Id for this product", http.StatusNotFound)
		return
	}

	json, err := json.MarshalIndent(review, "", "  ")
	if err != nil {
		http.Error(w, "could not marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func productReviewAddHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)

	prodId, err := strconv.ParseUint(vars["productid"], 10, 32)
	if err != nil {
		http.Error(w, "could not determine Id", http.StatusBadRequest)
		return
	}

	var review Review
	err = json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		http.Error(w, "could not unmarshal JSON", http.StatusBadRequest)
		return
	}

	AddReview(uint32(prodId), review)
}

func productReviewEditHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)

	prodId, err := strconv.ParseUint(vars["productid"], 10, 32)
	if err != nil {
		http.Error(w, "could not determine Id", http.StatusBadRequest)
		return
	}

	revId, err := strconv.ParseUint(vars["reviewid"], 10, 32)
	if err != nil {
		http.Error(w, "could not determine Id", http.StatusBadRequest)
		return
	}

	var review Review
	err = json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		http.Error(w, "could not unmarshal JSON", http.StatusBadRequest)
		return
	}

	if review.Id != uint32(revId) || review.ProductId != uint32(prodId) {
		http.Error(w, "could not find review for this product or with this Id", http.StatusBadRequest)
		return
	}

	err = EditReview(review)
	if err != nil {
		http.Error(w, "could not edit review", http.StatusInternalServerError)
		return
	}
}

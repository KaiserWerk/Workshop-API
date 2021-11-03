package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/*

 */

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	// no middleware
	router.HandleFunc("/api/v2/authenticate", nil).Methods(http.MethodGet)

	v1Router := router.PathPrefix("/api/v1").Subrouter()
	v1Router.Use(authV1)
	// Produkte
	v1Router.HandleFunc("/product/getall", productGetAllHandler).Methods(http.MethodGet)
	v1Router.HandleFunc("/product/{id}/get", productGetHandler).Methods(http.MethodGet)
	v1Router.HandleFunc("/product/add", productAddHandler).Methods(http.MethodPost)
	v1Router.HandleFunc("/product/edit", productEditHandler).Methods(http.MethodPut)
	v1Router.HandleFunc("/product/{id}/remove", productRemoveHandler).Methods(http.MethodDelete)

	v2Router := router.PathPrefix("/api/v2").Subrouter()
	v2Router.Use(authV2)
	// Produkte
	v2Router.HandleFunc("/product/getall", nil).Methods(http.MethodGet)
	v2Router.HandleFunc("/product/{id}/get", nil).Methods(http.MethodGet)
	v2Router.HandleFunc("/product/add", nil).Methods(http.MethodPost)
	v2Router.HandleFunc("/product/edit", nil).Methods(http.MethodPut)
	v2Router.HandleFunc("/product/{id}/delete", nil).Methods(http.MethodDelete)
	// Reviews
	v2Router.HandleFunc("/product/{productid}/review/getall", nil).Methods(http.MethodGet)
	v2Router.HandleFunc("/product/{productid}/review/{reviewid}/get", nil).Methods(http.MethodGet)
	v2Router.HandleFunc("/product/{productid}/review/add", nil).Methods(http.MethodPost)
	v2Router.HandleFunc("/product/{productid}/review/{reviewid}/edit", nil).Methods(http.MethodPut)

	srv := http.Server{
		Addr:         ":6789",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	fmt.Println("starting up server at port 6789...")
	log.Fatal(srv.ListenAndServe())
}

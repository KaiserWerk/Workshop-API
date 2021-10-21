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

	v1Router := router.PathPrefix("/api/v1").Subrouter()
	// Produkte
	v1Router.HandleFunc("/product/getall", nil).Methods(http.MethodGet)
	v1Router.HandleFunc("/product/{id}/get", nil).Methods(http.MethodGet)
	v1Router.HandleFunc("/product/add", nil).Methods(http.MethodPost)
	v1Router.HandleFunc("/product/{id}/edit", nil).Methods(http.MethodPut)
	v1Router.HandleFunc("/product/{id}/edit", nil).Methods(http.MethodDelete)

	v2Router := router.PathPrefix("/api/v2").Subrouter()
	// Authentifizierung
	v2Router.HandleFunc("/authenticate", nil).Methods(http.MethodGet)
	// Produkte
	v2Router.HandleFunc("/product/getall", nil).Methods(http.MethodGet)
	v2Router.HandleFunc("/product/{id}/get", nil).Methods(http.MethodGet)
	v2Router.HandleFunc("/product/add", nil).Methods(http.MethodPost)
	v2Router.HandleFunc("/product/{id}/edit", nil).Methods(http.MethodPut)
	v2Router.HandleFunc("/product/{id}/edit", nil).Methods(http.MethodDelete)
	// Produktbilder
	v2Router.HandleFunc("/image/{productid}/getall", nil).Methods(http.MethodGet)
	v2Router.HandleFunc("/image/{productid}/add", nil).Methods(http.MethodPost)
	v2Router.HandleFunc("/image/{productid}/replace", nil).Methods(http.MethodPut)
	v2Router.HandleFunc("/image/{productid}/remove", nil).Methods(http.MethodDelete)

	srv := http.Server{
		Addr:         ":6789",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	fmt.Println("starting up server at port 6789...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("server error: " + err.Error())
	}
}

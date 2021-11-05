package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var port int

func main() {
	flag.IntVar(&port, "Der Port für die Anwendung", 6789, "Gib den Port an, welches genutzt werden soll")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	// no middleware
	router.HandleFunc("/api/v2/authenticate", authenticateHandler).Methods(http.MethodGet)

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
	v2Router.HandleFunc("/product/getall", productGetAllHandler).Methods(http.MethodGet)
	v2Router.HandleFunc("/product/{id}/get", productGetHandler).Methods(http.MethodGet)
	v2Router.HandleFunc("/product/add", productAddHandler).Methods(http.MethodPost)
	v2Router.HandleFunc("/product/edit", productEditHandler).Methods(http.MethodPut)
	v2Router.HandleFunc("/product/{id}/delete", productRemoveHandler).Methods(http.MethodDelete)
	// Reviews
	v2Router.HandleFunc("/product/{productid}/review/getall", productReviewGetAllHandler).Methods(http.MethodGet)
	v2Router.HandleFunc("/product/{productid}/review/{reviewid}/get", productReviewGetHandler).Methods(http.MethodGet)
	v2Router.HandleFunc("/product/{productid}/review/add", productReviewAddHandler).Methods(http.MethodPost)
	v2Router.HandleFunc("/product/{productid}/review/{reviewid}/edit", productReviewEditHandler).Methods(http.MethodPut)

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	fmt.Printf("starting up server at port %d...\n", port)
	log.Fatal(srv.ListenAndServe())
}

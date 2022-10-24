package main

import (
	"github.com/gorilla/mux"
	"github.com/roneyhendrix/Golang-rest-sederhana/controllers/productcontroller"
	"github.com/roneyhendrix/Golang-rest-sederhana/models"
	"log"
	"net/http"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()
	r.HandleFunc("/api/products", productcontroller.ShowAllProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", productcontroller.ShowProduct).Methods("GET")
	r.HandleFunc("/api/product", productcontroller.CreateProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", productcontroller.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/product", productcontroller.DeleteProduct).Methods("DELETE")

	log.Println("Server started on port 9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}

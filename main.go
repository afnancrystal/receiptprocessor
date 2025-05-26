package main

import (
	"log"
	"net/http"
	"receiptprocessor/handlers"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("🚀 Receipt Processor running at http://localhost:8080")

	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

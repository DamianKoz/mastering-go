package main

import (
	"log"
	"net/http"
)

func main() {
	server := &VehicleServer{}
	log.Fatal(http.ListenAndServe(":8000", server))
}

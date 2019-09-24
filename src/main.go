package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type flower struct {
	Color string
}

// Get Env Vars
var color = os.Getenv("COLOR")

func main() {
	port := "80"
	http.HandleFunc("/", handler)
	log.Println("Listening on: " + port)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling request - color:[%s]", color)
	time.Sleep(1 * time.Second)

	// Format json response
	msg := flower{color}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

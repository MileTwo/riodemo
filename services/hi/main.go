package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type hiMessage struct {
	Message string
	Color   string
}

// Get Env Vars
var color = "Green"

func main() {
	port := "80"
	http.HandleFunc("/", handler)
	log.Println("Listening on: " + port)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling request with color:[%s]", color)
	time.Sleep(1 * time.Second)

	// Format json response
	msg := hiMessage{"hi", color}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

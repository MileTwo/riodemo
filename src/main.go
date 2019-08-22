package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type flower struct {
	Variety string
	Color   string
}

// Get Env Vars
var variety = os.Getenv("VARIETY")
var color = os.Getenv("COLOR")

func main() {
	port := "8080"
	http.HandleFunc("/", handler)
	log.Println("Listening on: " + port)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling request - variety: [%s], color:[%s]", variety, color)
	time.Sleep(1 * time.Second)

	// Format json response
	msg := flower{variety, color}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

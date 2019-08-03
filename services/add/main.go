package main

import (
    "log"
		"net/http"
		"fmt"
		"strconv"
		"encoding/json"
)
type addMessage struct {
	X int
	Y int
	Sum int
}

func main() {
		http.HandleFunc("/", handler)
		log.Println("Listening on 8080")
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	// Get parms
	x,_ := strconv.Atoi(r.URL.Query()["x"][0])
	y,_ := strconv.Atoi(r.URL.Query()["y"][0])

	// The algorithm 
	sum := x+y	
	fmt.Printf("computing %s + %s = %s\n",strconv.Itoa(x),strconv.Itoa(y),strconv.Itoa(sum))
	
	// Format json response
	msg := addMessage{x,y,sum}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
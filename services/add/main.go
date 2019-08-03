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
	port := "8080"
	http.HandleFunc("/", handler)
	log.Println("Listening on: " + port)
  http.ListenAndServe(":"+ port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	// Get parms
	x,_ := strconv.Atoi(r.URL.Query()["x"][0])
	y,_ := strconv.Atoi(r.URL.Query()["y"][0])

	// The algorithm 
	sum := x+y	
	fmt.Printf("computed %d + %d = %d\n",x,y,sum)
	
	// Format json response
	msg := addMessage{x,y,sum}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
package main

import (
    "log"
		"net/http"
		"fmt"
		"strconv"
		"encoding/json"
		"io/ioutil"
		"os"
)
type powerMessage struct {
	X int
	Y int
	Power int
}
type multMessage struct {
	X int
	Y int
	Product int
}

var MULT_URL string = os.Getenv("MULT_URL")

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
	// We could just do this 'power := math.Pow(float64(x),float64(y))' but what fun would that be?
	power := x
	for i := 1; i < y; i++ {
		power = mult(power,x)
	}
	fmt.Printf("computed %d^%d = %d\n",x,y,power)


	// Format json response
	msg := powerMessage{x,y,power}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func mult(x int, y int) (product int) {
	url := fmt.Sprintf("%s?x=%d&y=%d",MULT_URL,x,y)
	//log.Println("Call mult service at: " + url)
	resp, err := http.Get(url)
	if err != nil {log.Fatalln(err)}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	msg := multMessage{}
	jsonErr := json.Unmarshal(body, &msg)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	product = msg.Product
	return
}
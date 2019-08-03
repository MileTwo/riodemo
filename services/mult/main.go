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
type multMessage struct {
	X int
	Y int
	Product int
}
type addMessage struct {
	X int
	Y int
	Sum int
}

var ADD_URL string = os.Getenv("ADD_URL")

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
	// We could just do this 'product := x*y' but what fun would that be?
	product := y
	for i := 1; i < x; i++ {
		product = add(product,y)
	}
	fmt.Printf("computed %d x %d = %d\n",x,y,product)


	// Format json response
	msg := multMessage{x,y,product}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func add(x int, y int) (sum int) {
	url := fmt.Sprintf("%s?x=%d&y=%d",ADD_URL,x,y)
	//log.Println("Call add service at: " + url)
	resp, err := http.Get(url)
	if err != nil {log.Fatalln(err)}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	msg := addMessage{}
	jsonErr := json.Unmarshal(body, &msg)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	sum = msg.Sum
	return
}
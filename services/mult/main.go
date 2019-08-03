package main

import (
    "log"
		"net/http"
		"fmt"
		"strconv"
		"encoding/json"
		"io/ioutil"
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

func main() {
		http.HandleFunc("/", handler)
		log.Println("Listening on 8081")
    http.ListenAndServe(":8081", nil)
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
	fmt.Printf("computing %s x %s = %s\n",strconv.Itoa(x),strconv.Itoa(y),strconv.Itoa(product))


	// Format json response
	msg := multMessage{x,y,product}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func add(x int, y int) (sum int) {
	url := fmt.Sprintf("http://localhost:8080?x=%d&y=%d", x,y)
	resp, err := http.Get(url)
	if err != nil {log.Fatalln(err)}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	msg := addMessage{}
	jsonErr := json.Unmarshal(body, &msg)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Printf("The sum is: %d\n", msg.Sum)
	sum = msg.Sum
	return
}
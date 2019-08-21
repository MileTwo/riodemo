package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type multMessage struct {
	X       int
	Y       int
	Product int
}
type addMessage struct {
	X   int
	Y   int
	Sum int
}

var addURL = os.Getenv("ADD_URL")

func main() {
	port := "80"
	http.HandleFunc("/", handler)
	log.Println("Listening on: " + port)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Get parms
	x, _ := strconv.Atoi(r.URL.Query()["x"][0])
	y, _ := strconv.Atoi(r.URL.Query()["y"][0])

	// The algorithm
	// We could just do this 'product := x*y' but what fun would that be?
	product := y
	for i := 1; i < x; i++ {
		product = add(product, y)
	}
	log.Printf("computed %d x %d = %d\n", x, y, product)

	// Format json response
	msg := multMessage{x, y, product}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func add(x int, y int) (sum int) {
	url := fmt.Sprintf("%s?x=%d&y=%d", addURL, x, y)
	log.Println("Call add service at: " + url)

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
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

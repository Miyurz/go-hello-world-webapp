package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

func getMessage() (*Message, error) {
	greeterServiceHostname := "greeter-service"
	url := fmt.Sprintf("http://%s/", greeterServiceHostname)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var message Message
	json.Unmarshal(responseData, &message)
	return &message, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	message, err := getMessage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s (greeter-service: %s)", message.Message, message.Version)
}

func main() {
	fmt.Print("Starting 'hello-world-service'...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

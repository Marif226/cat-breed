package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Marif226/cat-breed/internal/model"
)

func main() {
	// Get the list of cat breeds from the API
	response, err := http.Get("https://catfact.ninja/breeds")
	if err != nil {
		log.Fatal("error during calling the API: ", err)
	}

	defer response.Body.Close()

	// Decode the JSON response
	var apiResp model.APIResponse
	err = json.NewDecoder(response.Body).Decode(&apiResp)
	if err != nil {
		log.Fatal("error during decoding response: ", err)
	}

	breeds := apiResp.Data

	file, err := os.Create("data/out.json")
	if err != nil {
		log.Fatal("error creating file: ", err)
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(breeds)
	if err != nil {
		log.Fatal("error writing result to file: ", err)
	}

	log.Println("successfully wrote the JSON data to out.json.")
}
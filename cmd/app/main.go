package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sort"
	"github.com/Marif226/cat-breed/internal/model"
	"github.com/Marif226/cat-breed/internal/utils"
)

func main() {
	// Get the list of cat breeds from the API
	response, err := http.Get("https://catfact.ninja/breeds")
	if err != nil {
		log.Fatalf("error during calling the API: %v", err)
	}

	defer response.Body.Close()

	// Decode the JSON response
	var apiResp model.APIResponse
	err = json.NewDecoder(response.Body).Decode(&apiResp)
	if err != nil {
		log.Fatalf("error during decoding response: %v", err)
	}

	breeds := apiResp.Data

	// Sort data by breed name length
	sort.Slice(breeds, func(i, j int) bool {
		return len(breeds[i].Name) < len(breeds[j].Name)
	})

	// Group data by country
	countryBreed := utils.GroupBreedByCountry(breeds)
	

	// Create directory for result
	err = os.MkdirAll("data", os.ModePerm)
	if err != nil {
		log.Fatalf("error creating directory: %v", err)
	}
	
	// Create result file
	file, err := os.Create("data/out.json")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}

	defer file.Close()

	// Marshal data with indentation
	data, err := json.MarshalIndent(countryBreed, "", "\t")
    if err != nil {
        log.Fatalf("error marshalling data: %v", err)
    }

	// Write data to file
	_, err = file.Write(data)
	if err != nil {
		log.Fatalf("error writing result to file: %v", err)
	}

	log.Println("successfully wrote the JSON data to out.json.")
}
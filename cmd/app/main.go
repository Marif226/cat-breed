package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sort"

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

	// Sort data by breed name length
	sort.Slice(breeds, func(i, j int) bool {
		return len(breeds[i].Name) < len(breeds[j].Name)
	})


	// Group data by country
	countryBreed := make(map[string]*model.BreedList)
	for _, b := range breeds {
		if _, ok := countryBreed[b.Country]; ok {
			countryBreed[b.Country].Data = append(countryBreed[b.Country].Data, b)
		} else {
			breed := model.BreedList{}
			breed.Country = b.Country
			breed.Data = append(breed.Data, b)
			countryBreed[b.Country] = &breed
		}
	}
	
	// Create result file
	file, err := os.Create("data/out.json")
	if err != nil {
		log.Fatal("error creating file: ", err)
	}

	defer file.Close()

	// Marshal data with indentation
	data, err := json.MarshalIndent(countryBreed, "", "\t")
    if err != nil {
        log.Fatal("error marshalling data:", err)
    }

	// Write data to file
	_, err = file.Write(data)
	if err != nil {
		log.Fatal("error writing result to file: ", err)
	}

	log.Println("successfully wrote the JSON data to out.json.")
}
package utils

import "github.com/Marif226/cat-breed/internal/model"

func GroupBreedByCountry(breeds []model.Breed) (map[string]*model.BreedList) {
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

	return countryBreed
}
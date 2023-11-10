package model

type Breed struct {
	Name 	string `json:"breed"`
	Country string `json:"country"`
}

type BreedList struct {
	Country	string	`json:"country"`
	Data	[]Breed	`json:"data"`
}

type APIResponse struct {
	CurrentPage	int		`json:"current_page"`
	Data		[]Breed	`json:"data"`
}
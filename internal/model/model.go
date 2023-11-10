package model

type breed struct {
	Name 	string `json:"breed"`
	Country string `json:"country"`
}

type APIResponse struct {
	CurrentPage	int		`json:"current_page"`
	Data		[]breed	`json:"data"`
}

type BreedsByCountry map[string][]string
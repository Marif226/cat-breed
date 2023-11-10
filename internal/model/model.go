package model

type breed struct {
	Name 	string `json:"breed"`
	Country string `json:"country"`
}

type BreedList struct {
	Country	string	`json:"country"`
	Data	[]breed	`json:"data"`
}

type APIResponse struct {
	CurrentPage	int		`json:"current_page"`
	Data		[]breed	`json:"data"`
}
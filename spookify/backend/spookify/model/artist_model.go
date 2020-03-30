package model

type Artist struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	Debut string `json:"debut"`
	ImageUrl string `json:"imageUrl"`
	Category string `json:"category"`
}

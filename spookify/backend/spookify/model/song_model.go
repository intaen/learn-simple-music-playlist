package model

type Song struct {
	ID int `json:"id,omitempty"`
	ArtistID int
	GenreID int
	Title string
	ImageUrl string
}

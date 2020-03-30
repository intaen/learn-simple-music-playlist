package artist

import "spookify/model"

type ArtistUsecase interface {
	GetAll() (*[]model.Artist, error)
	GetByID(id int) (*model.Artist, error)
	Insert(artist *model.Artist) error
	Update(id int, artist *model.Artist) error
	Delete(id int) error
	ArtistByGenre(id int) (*[]model.Artist, error)
}

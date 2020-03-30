package song

import "spookify/model"

type SongUsecase interface {
	GetAll() (*[]model.Song, error)
	GetByID(id int) (*model.Song, error)
	SongByArtist(id int) (*[]model.Song, error)
	Insert(song *model.Song) error
	Update(id int, genre *model.Song) error
	Delete(id int) error
}
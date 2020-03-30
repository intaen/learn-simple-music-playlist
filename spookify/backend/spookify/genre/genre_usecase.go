package genre

import "spookify/model"

type GenreUsecase interface {
	GetAll() (*[]model.Genre, error)
	GetByID(id int) (*model.Genre, error)
	Insert(genre *model.Genre) error
	Update(id int, genre *model.Genre) error
	Delete(id int) error
}
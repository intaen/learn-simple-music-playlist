package usecase

import (
	"spookify/genre"
	"spookify/model"
)

type GenreUsecaseImpl struct {
	genreRepo genre.GenreRepo
}

func (g *GenreUsecaseImpl) GetAll() (*[]model.Genre, error) {
	return g.genreRepo.GetAll()
}

func (g *GenreUsecaseImpl) GetByID(id int) (*model.Genre, error) {
	return g.genreRepo.GetByID(id)
}

func (g *GenreUsecaseImpl) Insert(genre *model.Genre) error {
	return g.genreRepo.Insert(genre)
}

func (g *GenreUsecaseImpl) Update(id int, genre *model.Genre) error {
	return g.genreRepo.Update(id, genre)
}

func (g *GenreUsecaseImpl) Delete(id int) error {
	return g.genreRepo.Delete(id)
}

func CreateGenreUsecaseImpl (genreRepo genre.GenreRepo) genre.GenreUsecase {
	return &GenreUsecaseImpl{genreRepo}
}
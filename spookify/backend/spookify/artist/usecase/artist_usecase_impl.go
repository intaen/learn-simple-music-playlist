package usecase

import (
	"spookify/artist"
	"spookify/model"
)

type ArtistUsecaseImpl struct {
	artistRepo artist.ArtistRepo
}

func (a *ArtistUsecaseImpl) GetAll() (*[]model.Artist, error) {
	return a.artistRepo.GetAll()
}

func (a *ArtistUsecaseImpl) GetByID(id int) (*model.Artist, error) {
	return a.artistRepo.GetByID(id)
}

func (a *ArtistUsecaseImpl) Insert(artist *model.Artist) error {
	return a.artistRepo.Insert(artist)
}

func (a *ArtistUsecaseImpl) Update(id int, artist *model.Artist) error {
	return a.artistRepo.Update(id, artist)
}

func (a *ArtistUsecaseImpl) Delete(id int) error {
	return a.artistRepo.Delete(id)
}

func (a *ArtistUsecaseImpl) ArtistByGenre(id int) (*[]model.Artist, error) {
	return a.artistRepo.ArtistByGenre(id)
}

func CreateArtistUsecaseImpl (artistRepo artist.ArtistRepo) artist.ArtistUsecase {
	return &ArtistUsecaseImpl{artistRepo}
}
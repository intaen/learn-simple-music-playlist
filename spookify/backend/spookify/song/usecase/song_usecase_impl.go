package usecase

import (
	"spookify/song"
	"spookify/model"
)

type SongUsecaseImpl struct {
	songRepo song.SongRepo
}

func (s *SongUsecaseImpl) GetAll() (*[]model.Song, error) {
	return s.songRepo.GetAll()
}

func (s *SongUsecaseImpl) GetByID(id int) (*model.Song, error) {
	return s.songRepo.GetByID(id)
}

func (s *SongUsecaseImpl) SongByArtist(id int) (*[]model.Song, error) {
	return s.songRepo.SongByArtist(id)
}

func (s *SongUsecaseImpl) Insert(song *model.Song) error {
	return s.songRepo.Insert(song)
}

func (s *SongUsecaseImpl) Update(id int, song *model.Song) error {
	return s.songRepo.Update(id, song)
}

func (s *SongUsecaseImpl) Delete(id int) error {
	return s.songRepo.Delete(id)
}

func CreateSongUsecaseImpl (songRepo song.SongRepo) song.SongUsecase {
	return &SongUsecaseImpl{songRepo}
}
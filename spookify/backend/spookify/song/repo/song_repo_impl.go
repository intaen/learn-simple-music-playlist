package repo

import (
	"fmt"

	"database/sql"

	"spookify/song"
	"spookify/model"
)

type SongRepoImpl struct {
	db *sql.DB
}

func (s *SongRepoImpl) GetAll() (*[]model.Song, error) {
	var song = model.Song{}
	var arrSong []model.Song

	query := "SELECT id, artist_id, genre_id, title, image_url FROM songs"
	rows, err := s.db.Query(query)
	if err != nil {
		fmt.Printf("[SongRepoImpl.GetDataByUser] Error when select: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&song.ID, &song.ArtistID, &song.GenreID, &song.Title, &song.ImageUrl)
		if err != nil {
			fmt.Errorf("[SongRepoImpl.GetDataByUser] Error when scanning: %w", err)
		}
		arrSong = append(arrSong, song)
	}
	return &arrSong, nil
}

func (s *SongRepoImpl) GetByID(id int) (*model.Song, error) {
	var song model.Song

	query := "SELECT id, artist_id, genre_id, title, image_url FROM songs WHERE artist_id = ?"
	s.db.QueryRow(query, id).Scan(&song.ID, &song.ArtistID, &song.GenreID, &song.Title, &song.ImageUrl)
	return &song, nil
}

func (s *SongRepoImpl) SongByArtist(id int) (*[]model.Song, error) {
	var song = model.Song{}
	var arrSong []model.Song

	query := "SELECT songs.title, songs.image_url FROM songs JOIN artists ON artists.id = songs.artist_id where artist_id = ?"

	rows, err := s.db.Query(query, id)
	if err != nil {
		fmt.Printf("[SongRepoImpl.SongByArtist] Error when select: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&song.Title, &song.ImageUrl)
		if err != nil {
			fmt.Errorf("[ArtistRepoImpl.GetDataByUser] Error when scanning: %w", err)
		}
		arrSong = append(arrSong, song)
	}
	return &arrSong, nil
}

func (s *SongRepoImpl) Insert(song *model.Song) error {
	query := "INSERT INTO songs (artist_id, genre_id, title, image_url) VALUES (?, ?, ?, ?)"

	_, err := s.db.Exec(query, &song.ArtistID, &song.GenreID, &song.Title, &song.ImageUrl)
	if err != nil {
		fmt.Errorf("[SongRepoImpl.Insert] Error when execute db: %w", err)
	}
	return nil
}

func (s *SongRepoImpl) Update(id int, song *model.Song) error {
	query := "UPDATE songs SET artist_id=?, genre_id=?, title=?, image_url=? WHERE id=?"

	_, err := s.db.Exec(query, &song.ArtistID, &song.GenreID, &song.Title, &song.ImageUrl, id)
	if err != nil{
		fmt.Errorf("[SongRepoImpl.Update] Error when execute query: %w", err)
	}
	return nil
}

func (s *SongRepoImpl) Delete(id int) error {
	query := "DELETE FROM songs WHERE id=?"

	_, err:= s.db.Exec(query, id)
	if err != nil {
		fmt.Errorf("[SongRepoImpl.Delete] Error when execute query: %w", err)
	}
	return nil
}

func CreateSongRepoImpl (db *sql.DB) song.SongRepo {
	return &SongRepoImpl{db}
}
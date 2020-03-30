package repo

import (
	"fmt"

	"database/sql"

	"spookify/artist"
	"spookify/model"
)

type ArtistRepoImpl struct {
	db *sql.DB
}

func (a *ArtistRepoImpl) GetAll() (*[]model.Artist, error) {
	var artist = model.Artist{}
	var arrArtist []model.Artist

	query := "SELECT id, name, debut, image_url, category FROM artists"
	rows, err := a.db.Query(query)
	if err != nil {
		fmt.Printf("[ArtistRepoImpl.GetDataByUser] Error when select all artists: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&artist.ID, &artist.Name, &artist.Debut, &artist.ImageUrl, &artist.Category)
		if err != nil {
			fmt.Errorf("[ArtistRepoImpl.GetDataByUser] Error when scanning: %w", err)
		}
		arrArtist = append(arrArtist, artist)
	}
	return &arrArtist, nil
}

func (a *ArtistRepoImpl) GetByID(id int) (*model.Artist, error) {
	var artist model.Artist

	query := "SELECT id, name, debut, image_url, category FROM artists WHERE id = ?"
	a.db.QueryRow(query, id).Scan(&artist.ID, &artist.Name, &artist.Debut, &artist.ImageUrl, &artist.Category)
	return &artist, nil
}

func (a *ArtistRepoImpl) Insert(artist *model.Artist) error {
	query := "INSERT INTO artists (name, debut, image_url, category) VALUES (?, ?, ?, ?)"

	_, err := a.db.Exec(query, &artist.Name, &artist.Debut, &artist.ImageUrl, &artist.Category)
	if err != nil {
		fmt.Errorf("[ArtistRepoImpl.Insert] Error when execute db: %w", err)
	}
	return nil
}

func (a *ArtistRepoImpl) Update(id int, artist *model.Artist) error {
	query := "UPDATE artists SET name=?, debut=?, image_url=?, category=? WHERE id=?"

	_, err := a.db.Exec(query, &artist.Name, &artist.Debut, &artist.ImageUrl, &artist.Category, id)
	if err != nil{
		fmt.Errorf("[ArtistRepoImpl.Update] Errorn when execute query: %w", err)
	}
	return nil
}

func (a *ArtistRepoImpl) Delete(id int) error {
	query := "DELETE FROM artists WHERE id=?"

	_, err:= a.db.Exec(query, id)
	if err != nil {
		fmt.Errorf("[ArtistRepoImpl.Delete] Errorn when execute query: %w", err)
	}
	return nil
}

func (a *ArtistRepoImpl) ArtistByGenre(id int) (*[]model.Artist, error) {
	var artist = model.Artist{}
	var arrArtist []model.Artist

	query := "SELECT artists.id, artists.name,artists.debut, artists.image_url,	artists.category FROM artists JOIN songs ON artists.id = songs.artist_id WHERE genre_id = ? GROUP BY artists.name"

	rows, err := a.db.Query(query, id)
	if err != nil {
		fmt.Printf("[ArtistRepoImpl.ArtistByGenre] Error when select: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&artist.ID, &artist.Name, &artist.Debut, &artist.ImageUrl, &artist.Category)
		if err != nil {
			fmt.Errorf("[ArtistRepoImpl.GetDataByUser] Error when scanning: %w", err)
		}
		arrArtist = append(arrArtist, artist)
	}
	return &arrArtist, nil
}

func CreateArtistRepoImpl (db *sql.DB) artist.ArtistRepo {
	return &ArtistRepoImpl{db}
}
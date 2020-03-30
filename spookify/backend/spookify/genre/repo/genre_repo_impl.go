package repo

import (
	"fmt"

	"database/sql"

	"spookify/genre"
	"spookify/model"
)

type GenreRepoImpl struct {
	db *sql.DB
}

func (g *GenreRepoImpl) GetAll() (*[]model.Genre, error) {
	var genre = model.Genre{}
	var arrGenre []model.Genre

	query := "SELECT id, type FROM genres"
	rows, err := g.db.Query(query)
	if err != nil {
		fmt.Printf("[GenreRepoImpl.GetDataByUser] Error when select: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&genre.ID, &genre.Type)
		if err != nil {
			fmt.Errorf("[GenreRepoImpl.GetDataByUser] Error when scanning: %w", err)
		}
		arrGenre = append(arrGenre, genre)
	}
	return &arrGenre, nil
}

func (g *GenreRepoImpl) GetByID(id int) (*model.Genre, error) {
	var genre model.Genre

	query := "SELECT id, type FROM genres WHERE id = ?"
	g.db.QueryRow(query, id).Scan(&genre.ID, &genre.Type)
	return &genre, nil
}

func (g *GenreRepoImpl) Insert(genre *model.Genre) error {
	query := "INSERT INTO genres (type) VALUES (?)"

	_, err := g.db.Exec(query, genre.Type)
	if err != nil {
		fmt.Errorf("[GenreRepoImpl.Insert] Error when execute db: %w", err)
	}
	return nil
}

func (g *GenreRepoImpl) Update(id int, genre *model.Genre) error {
	query := "UPDATE genres SET type=? WHERE id=?"

	_, err := g.db.Exec(query, &genre.Type, id)
	if err != nil{
		fmt.Errorf("[GenreRepoImpl.Update] Error when execute query: %w", err)
	}
	return nil
}

func (g *GenreRepoImpl) Delete(id int) error {
	query := "DELETE FROM genres WHERE id=?"

	_, err:= g.db.Exec(query, id)
	if err != nil {
		fmt.Errorf("[GenreRepoImpl.Delete] Errorn when execute query: %w", err)
	}
	return nil
}

func CreateGenreRepoImpl (db *sql.DB) genre.GenreRepo {
	return &GenreRepoImpl{db}
}
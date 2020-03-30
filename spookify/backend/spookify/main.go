package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"spookify/middleware"

	ah "spookify/artist/handler"
	au "spookify/artist/usecase"
	ar "spookify/artist/repo"

	gh "spookify/genre/handler"
	gu "spookify/genre/usecase"
	gr "spookify/genre/repo"

	sh "spookify/song/handler"
	su "spookify/song/usecase"
	sr "spookify/song/repo"
)

func main() {
	port := "8081"
	connect := "root:admin@tcp(127.0.0.1:3306)/spookify"

	db, err := sql.Open("mysql", connect)
	if err != nil {
		log.Fatal("Can't connect to database " + connect + ":" + err.Error())
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

	artistRepo := ar.CreateArtistRepoImpl(db)
	artistUsecase := au.CreateArtistUsecaseImpl(artistRepo)
	ah.CreateArtistHandler(router, artistUsecase)

	genreRepo := gr.CreateGenreRepoImpl(db)
	genreUsecase := gu.CreateGenreUsecaseImpl(genreRepo)
	gh.CreateGenreHandler(router, genreUsecase)

	songRepo := sr.CreateSongRepoImpl(db)
	songUsecase := su.CreateSongUsecaseImpl(songRepo)
	sh.CreateSongHandler(router, songUsecase)

	router.Use(middleware.Logger)

	fmt.Println("Starting Web Server at Port: " + port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

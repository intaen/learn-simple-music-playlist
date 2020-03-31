# Music Playlist
##### @intanmarsjaf
-----------------------
### Mobile Apps Simple Music Playlist using Kotlin and Golang
So, I get task to make a simple apps like spotify (thats why I name the apps with spookify), and I make the simple apps, still no music player, just get the list of artist, genre, and song.
I use two programming language to build this, golang for backend and kotlin for frontend. Stack in this apps are:
#### Golang:
+ RestAPI
+ MySql
+ Gorilla Mux
#### Kotlin:
+ Volley
+ Custom Array Adapter
#### Assets:
+ fonts.google.com
#### To run this apps, all you have to do:
+ Create database (I set up my database with name spookify)
+ Create table artists, genres, songs
```sh
CREATE TABLE artists (
	id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(30),
    debut DATE,
    image_url longtext,
    category VARCHAR(30)
);

CREATE TABLE genres (
	id INT PRIMARY KEY AUTO_INCREMENT,
    type VARCHAR(30)
);

CREATE TABLE songs (
	id INT PRIMARY KEY AUTO_INCREMENT,
    artist_id INT,
    genre_id INT,
    title VARCHAR(45),
    image_url longtext,
    FOREIGN KEY (artist_id)
		REFERENCES artists(id),
	FOREIGN KEY (genre_id)
		REFERENCES genres(id)
);
```
+ Run main.go file and the server will start in localhost:8081
```sh
go run main.go
```
+ Then, open android studio and start emulator, klik l>, wait for awhile...
+ Because the database is empty, we have to add user
+ So, this is API for Spookify:
#### Insert Artist
```sh
http://localhost:8081/artist
```
```sh
{
	"name": "artist_name",
	"debut": "debut_date",
	"imageUrl": "http://link_img_of_artist",
	"category": "category_artist, ex: solo"
}
```
#### GetAll Artist
```sh
http://localhost:8081/artists
```
#### GetArtistByGenre
```sh
http://localhost:8081/artist/genre/{id}
```
#### Insert Genre
```sh
http://localhost:8081/genre
```
```sh
{
    "type": "genre_artist"
}
```
#### GetAll Genre
```sh
http://localhost:8081/genres
```
#### Insert Song
```sh
http://localhost:8081/song
```
```sh
{
    "artist_id": 1,
	"genre_id": 1,
	"title": "song_title",
	"image_url": "http://link_of_song_img"
}
```
#### GetAll Song
```sh
http://localhost:8081/songs
```
#### GetSongByArtist
```sh
http://localhost:8081/song/artist/{id}
```

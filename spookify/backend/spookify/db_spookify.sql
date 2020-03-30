use spookify;

SELECT * FROM artists;
DELETE FROM artists;

SELECT * FROM genres;
DELETE FROM genres;

SELECT * FROM songs;
DELETE FROM songs;

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
	
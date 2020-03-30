SELECT artists.id, artists.name,artists.debut, artists.image_url,
	artists.category
		FROM artists JOIN songs ON artists.id = songs.artist_id where genre_id = 1 group by artists.name;
        
SELECT * FROM songs;

SELECT songs.title, songs.image_url FROM songs JOIN artists ON artists.id = songs.artist_id where artist_id = 1;
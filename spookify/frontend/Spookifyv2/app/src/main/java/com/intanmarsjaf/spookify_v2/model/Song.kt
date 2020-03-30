package com.intanmarsjaf.spookify_v2.model

class Song(var ArtistID: Int, var GenreID: Int, var Title: String, var ImageUrl: String) {
    override fun toString(): String {
        return "Song(ArtistID=$ArtistID, GenreID=$GenreID, Title='$Title')"
    }
}
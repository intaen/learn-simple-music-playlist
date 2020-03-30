package com.intanmarsjaf.spookify_v2.model

class Artist(val id: Int, val Name: String, var Debut: String, var ImageUrl: String, var Category: String) {
    override fun toString(): String {
        return "Artist(Name='$Name', Debut='$Debut', Category='$Category')"
    }
}
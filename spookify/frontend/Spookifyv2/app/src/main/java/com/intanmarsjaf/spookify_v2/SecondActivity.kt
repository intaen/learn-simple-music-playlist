package com.intanmarsjaf.spookify_v2

import android.content.Intent
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.view.View

class SecondActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_second)
    }

    fun genre(view: View){
        val newIntent = Intent(this, GenreActivity::class.java)
        startActivity(newIntent)
    }

    fun song(view: View){
        val newIntent = Intent(this, SongActivity::class.java)
        startActivity(newIntent)
    }

    fun artist(view: View){
        val newIntent = Intent(this, ArtistActivity::class.java)
        startActivity(newIntent)
    }
}

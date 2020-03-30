package com.intanmarsjaf.spookify_v2

import android.content.Intent
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.util.Log
import android.widget.ListView
import com.android.volley.Request
import com.android.volley.RequestQueue
import com.android.volley.Response
import com.android.volley.VolleyError
import com.android.volley.toolbox.StringRequest
import com.android.volley.toolbox.Volley
import com.intanmarsjaf.spookify_v2.adapter.GenreArrayAdapter
import com.intanmarsjaf.spookify_v2.model.Genre
import org.json.JSONArray
import org.json.JSONException

class GenreActivity : AppCompatActivity() {

    val activityName = "Genre"
    lateinit var listView: ListView
    var listGenre = mutableListOf<Genre>()
    lateinit var arrayAdapter: GenreArrayAdapter
    lateinit var requestQue: RequestQueue
    val url = "http://10.0.2.2:8081//genre"

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.genre_list)
        Log.i(activityName, "OnCreate() Called")
        listView = findViewById<ListView>(R.id.genre_list)

        arrayAdapter = GenreArrayAdapter(
            context = this,
            genreList = listGenre
        )
        requestQue = Volley.newRequestQueue(this)
        listView.adapter = arrayAdapter

        listView.setOnItemClickListener { _, _, position, _ ->
            startActivity(Intent(this,ArtistByGenre::class.java).apply {
                putExtra("id",listGenre[position].id.toString())
            })
        }
        fetchAll()
    }

    fun fetchAll() {
        val genreRequest = StringRequest(
            Request.Method.GET,
            url,
            Response.Listener { response -> resolveSucces(response) },
            Response.ErrorListener { error: VolleyError? ->
                Log.e("FETCH: FAIL: ", error.toString())
            })

        requestQue.add(genreRequest)
    }

    private fun resolveSucces(response: String?) {
        try {
            val arrayResponse = JSONArray(response)
            for (i in 0 until arrayResponse.length()) {
                val item = arrayResponse.getJSONObject(i)
                val genre = Genre(
                    item.getInt("id"),
                    item.getString("Type")
                )
                Log.i("GENRE $i", genre.toString())
                arrayAdapter.add(genre)
            }
        } catch (jsonEx: JSONException) {
            Log.e("PARSE: FAIL:", jsonEx.message)
        }
    }
}

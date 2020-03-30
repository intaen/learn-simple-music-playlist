package com.intanmarsjaf.spookify_v2

import android.content.Intent
import android.os.Bundle
import android.util.Log
import android.widget.ListView
import androidx.appcompat.app.AppCompatActivity
import com.android.volley.Request
import com.android.volley.RequestQueue
import com.android.volley.Response
import com.android.volley.VolleyError
import com.android.volley.toolbox.StringRequest
import com.android.volley.toolbox.Volley
import com.intanmarsjaf.spookify_v2.adapter.ArtistArrayAdapter
import com.intanmarsjaf.spookify_v2.model.Artist
import org.json.JSONArray
import org.json.JSONException

class ArtistByGenre: AppCompatActivity() {

    val activityName = "ArtistByGenre"
    lateinit var listView: ListView
    var listArtist = mutableListOf<Artist>()
    lateinit var arrayAdapter: ArtistArrayAdapter
    lateinit var requestQue: RequestQueue
    val url = "http://10.0.2.2:8081//artist/genre/"

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.artistbygenre_list)
        Log.i(activityName, "OnCreate() Called")

        listView = findViewById<ListView>(R.id.artistbygenre_list)
        arrayAdapter = ArtistArrayAdapter(
            context = this,
            artistList = listArtist
        )
        requestQue = Volley.newRequestQueue(this)
        listView.adapter = arrayAdapter

        listView.setOnItemClickListener { _, _, position, _ ->
            startActivity(Intent(this,SongByArtist::class.java).apply {
                putExtra("id",listArtist[position].id.toString())
            })
        }
        fetchAll()
    }

    fun fetchAll() {
        val artistRequest = StringRequest(
            Request.Method.GET,
            url+this.intent.getStringExtra("id")?.toString(),
            Response.Listener { response -> resolveSucces(response) },
            Response.ErrorListener { error: VolleyError? ->
                Log.e("FETCH: FAIL: ", error.toString())
            })

        requestQue.add(artistRequest)
    }

    private fun resolveSucces(response: String?) {
        try {
            val arrayResponse = JSONArray(response)
            for (i in 0 until arrayResponse.length()) {
                val item = arrayResponse.getJSONObject(i)
                val artist = Artist(
                    item.getInt("id"),
                    item.getString("Name"),
                    item.getString("Debut"),
                    item.getString("ImageUrl"),
                    item.getString("Category")
                )
                Log.i("ARTIST $i", artist.toString())
                arrayAdapter.add(artist)
            }
        } catch (jsonEx: JSONException) {
            Log.e("PARSE: FAIL:", jsonEx.message)
        }
    }
}
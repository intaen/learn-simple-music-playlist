package com.intanmarsjaf.spookify_v2

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
import com.intanmarsjaf.spookify_v2.adapter.ArtistArrayAdapter
import com.intanmarsjaf.spookify_v2.model.Artist
import org.json.JSONArray
import org.json.JSONException

class ArtistActivity : AppCompatActivity() {

    val activityName = "Artist"
    lateinit var listView: ListView
    var listArtist = mutableListOf<Artist>()
    lateinit var arrayAdapter: ArtistArrayAdapter
    lateinit var requestQue: RequestQueue
    val url = "http://10.0.2.2:8081/artist"

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.artist_list)

        Log.i(activityName, "OnCreate() Called")

        listView = findViewById<ListView>(R.id.artist_list)
        arrayAdapter = ArtistArrayAdapter(
            context = this,
            artistList = listArtist
        )
        requestQue = Volley.newRequestQueue(this)
        listView.adapter = arrayAdapter
        fetchAll()
    }

    fun fetchAll() {
        val artistRequest = StringRequest(
            Request.Method.GET,
            url,
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

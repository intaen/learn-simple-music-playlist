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
import com.intanmarsjaf.spookify_v2.adapter.SongArrayAdapter
import com.intanmarsjaf.spookify_v2.model.Song
import org.json.JSONArray
import org.json.JSONException

class SongByArtist : AppCompatActivity() {

    val activityName = "Song"
    lateinit var listView: ListView
    var listSong = mutableListOf<Song>()
    lateinit var arrayAdapter: SongArrayAdapter
    lateinit var requestQue: RequestQueue
    val url = "http://10.0.2.2:8081//song/artist/"

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.song_list)
        Log.i(activityName, "OnCreate() Called")

        listView = findViewById<ListView>(R.id.song_list)
        arrayAdapter = SongArrayAdapter(
            context = this,
            songList = listSong
        )
        requestQue = Volley.newRequestQueue(this)
        listView.adapter = arrayAdapter
        fetchAll()
    }

    fun fetchAll() {
        val songRequest = StringRequest(
            Request.Method.GET,
            url+this.intent.getStringExtra("id")?.toString(),
            Response.Listener { response -> resolveSucces(response) },
            Response.ErrorListener { error: VolleyError? ->
                Log.e("FETCH: FAIL: ", error.toString())
            })

        requestQue.add(songRequest)
    }

    private fun resolveSucces(response: String?) {
        try {
            val arrayResponse = JSONArray(response)
            for (i in 0 until arrayResponse.length()) {
                val item = arrayResponse.getJSONObject(i)
                val song = Song(
                    item.getInt("ArtistID"),
                    item.getInt("GenreID"),
                    item.getString("Title"),
                    item.getString("ImageUrl")
                )
                Log.i("SONG $i", song.toString())
                arrayAdapter.add(song)
            }
        } catch (jsonEx: JSONException) {
            Log.e("PARSE: FAIL:", jsonEx.message)
        }
    }
}

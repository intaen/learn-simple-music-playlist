package com.intanmarsjaf.spookify_v2.adapter

import android.content.Context
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ArrayAdapter
import android.widget.ImageView
import android.widget.TextView
import androidx.annotation.LayoutRes
import androidx.annotation.NonNull
import com.intanmarsjaf.spookify_v2.model.Song
import com.intanmarsjaf.spookify_v2.R
import com.squareup.picasso.Picasso

class SongArrayAdapter(@NonNull context: Context, @LayoutRes layoutRes: Int=0, var songList: MutableList<Song>): ArrayAdapter<Song>(context, layoutRes, songList) {
    override fun getView(position: Int, convertView: View?, parent: ViewGroup): View {
        var itemView = LayoutInflater.from(context).inflate(R.layout.activity_song,parent,false)
        val song = songList.get(position)
        itemView?.findViewById<TextView>(R.id.song_title)?.setText(song.Title)
        val imageView = itemView?.findViewById<ImageView>(R.id.imageView)
        Picasso.get().load(song.ImageUrl).placeholder(R.drawable.ic_account_circle_black_24dp).into(imageView)
        return itemView
    }
}
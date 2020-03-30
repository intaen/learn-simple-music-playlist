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
import com.intanmarsjaf.spookify_v2.R
import com.intanmarsjaf.spookify_v2.model.Artist
import com.squareup.picasso.Picasso

class ArtistArrayAdapter(@NonNull context: Context, @LayoutRes layoutRes: Int=0, var artistList: MutableList<Artist>): ArrayAdapter<Artist>(context, layoutRes, artistList) {
    override fun getView(position: Int, convertView: View?, parent: ViewGroup): View {
        var itemView = LayoutInflater.from(context).inflate(R.layout.activity_artist,parent,false)
        val artist = artistList.get(position)
        itemView?.findViewById<TextView>(R.id.artist_name)?.setText(artist.Name)
        itemView?.findViewById<TextView>(R.id.artist_debut)?.setText(artist.Debut)
        itemView?.findViewById<TextView>(R.id.artist_category)?.setText(artist.Category)
        val imageView = itemView?.findViewById<ImageView>(R.id.imageView)
        Picasso.get().load(artist.ImageUrl).placeholder(R.drawable.ic_account_circle_black_24dp).into(imageView)
        return itemView
    }
}
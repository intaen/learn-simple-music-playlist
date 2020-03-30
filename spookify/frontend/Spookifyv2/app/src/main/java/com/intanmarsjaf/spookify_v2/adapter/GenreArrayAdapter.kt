package com.intanmarsjaf.spookify_v2.adapter

import android.content.Context
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ArrayAdapter
import android.widget.TextView
import androidx.annotation.LayoutRes
import androidx.annotation.NonNull
import com.intanmarsjaf.spookify_v2.R
import com.intanmarsjaf.spookify_v2.model.Genre

class GenreArrayAdapter(@NonNull context: Context, @LayoutRes layoutRes: Int=0, var genreList: MutableList<Genre>): ArrayAdapter<Genre>(context, layoutRes, genreList) {
    override fun getView(position: Int, convertView: View?, parent: ViewGroup): View {
        var itemView = LayoutInflater.from(context).inflate(R.layout.activity_genre,parent,false)
        val genre = genreList.get(position)
        itemView?.findViewById<TextView>(R.id.genre_type)?.setText(genre.Type)
        return itemView
    }
}
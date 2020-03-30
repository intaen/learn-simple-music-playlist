package handler

import (
	"fmt"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/gorilla/mux"

	"spookify/song"
	"spookify/model"
	"spookify/response"

)

type SongHandler struct {
	songUsecase song.SongUsecase
}

func (s *SongHandler) getAllHandler(resp http.ResponseWriter, req *http.Request) {

	song, err := s.songUsecase.GetAll()
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Printf("[SongHandler.getAllHandler] Error when: %w", err)
		return
	}
	//response.HandleSuccess(resp, song)
	jsonData, err := json.Marshal(song)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[ShipHandler.handleSuccess] Error when do json marshalling for error handling: %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
	return
}

func (s *SongHandler) getByIdHandler(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		response.HandleError(resp, "ID must number")
		fmt.Printf("[SongHandler.getByIdHandler] Error when %w", err)
		return
	}

	song, err := s.songUsecase.GetByID(id)
	if err != nil {
		response.HandleError(resp, "ID must number")
		fmt.Printf("[SongHandler.getByIdHandler] Error when %w", err)
		return
	}
	//response.HandleSuccess(resp, song)
	jsonData, err := json.Marshal(song)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[ShipHandler.handleSuccess] Error when do json marshalling for error handling: %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
	return
}

func (s *SongHandler) getSongByArtistHandler(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		response.HandleError(resp, "ID must number")
		fmt.Printf("[ArtistHandler.getByIdHandler] Error when %w", err)
		return
	}

	song, err := s.songUsecase.SongByArtist(id)
	if err != nil {
		response.HandleError(resp, "ID must number")
		fmt.Printf("[ArtistHandler.getByIdHandler] Error when %w", err)
		return
	}
	//response.HandleSuccess(resp, artist)
	jsonData, err := json.Marshal(song)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[ShipHandler.handleSuccess] Error when do json marshalling for error handling: %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
	return
}

func (s *SongHandler) insertHandler(resp http.ResponseWriter, req *http.Request) {
	reqbody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[SongHandler.insertHandler] Error when reading body: " + err.Error())
		return
	}

	var song = model.Song{}
	err = json.Unmarshal(reqbody, &song)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[SongHandler.insertHandler] Error when reading body: " + err.Error())
		return
	}

	err = s.songUsecase.Insert(&song)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[SongHandler.insertHandler] Error when reading body: " + err.Error())
		return
	}
	response.HandleSuccess(resp, nil)
}

func (s *SongHandler) updateHandler(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[SongHandler.updateHandler] Error when convert: " + err.Error())
		return
	}

	data := model.Song{}
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[SongHandler.updateHandler] Error when reading: " + err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[SongHandler.updateHandler] Error when unmarshall: " + err.Error())
		return
	}

	song := s.songUsecase.Update(id, &data)
	response.HandleSuccess(resp, song)
}

func (s *SongHandler) deleteHandler(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[SongHandler.updateHandler] Error when: " + err.Error())
		return
	}

	song := s.songUsecase.Delete(id)
	response.HandleSuccess(resp, song)
}

func CreateSongHandler(r *mux.Router, songUsecase song.SongUsecase) {

	songHandler := SongHandler{songUsecase}

	r.HandleFunc("/song", songHandler.getAllHandler).Methods(http.MethodGet)
	s := r.PathPrefix("/song").Subrouter()
	s.HandleFunc("/{id}", songHandler.getByIdHandler).Methods(http.MethodGet)
	s.HandleFunc("/artist/{id}", songHandler.getSongByArtistHandler).Methods(http.MethodGet)
	s.HandleFunc("/add", songHandler.insertHandler).Methods(http.MethodPost)
	s.HandleFunc("/{id}", songHandler.updateHandler).Methods(http.MethodPut)
	s.HandleFunc("/{id}", songHandler.deleteHandler).Methods(http.MethodDelete)
}
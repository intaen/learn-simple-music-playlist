package handler

import (
	"fmt"
	"strconv"

	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/gorilla/mux"

	"spookify/artist"
	"spookify/model"
	"spookify/response"
)

type ArtistHandler struct {
	artistUsecase artist.ArtistUsecase
}

func (a *ArtistHandler) getAllHandler(resp http.ResponseWriter, req *http.Request) {

	artist, err := a.artistUsecase.GetAll()
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Printf("[ArtistHandler.getAllHandler] Error when: %w", err)
		return
	}
	//response.HandleSuccess(resp, artist)
	jsonData, err := json.Marshal(artist)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[ShipHandler.handleSuccess] Error when do json marshalling for error handling: %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
	return
}

func (a *ArtistHandler) getByIdHandler(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		response.HandleError(resp, "ID must number")
		fmt.Printf("[ArtistHandler.getByIdHandler] Error when %w", err)
		return
	}

	artist, err := a.artistUsecase.GetByID(id)
	if err != nil {
		response.HandleError(resp, "ID must number")
		fmt.Printf("[ArtistHandler.getByIdHandler] Error when %w", err)
		return
	}
	//response.HandleSuccess(resp, artist)
	jsonData, err := json.Marshal(artist)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[ShipHandler.handleSuccess] Error when do json marshalling for error handling: %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
	return
}

func (a* ArtistHandler) getArtistByGenreHandler(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		response.HandleError(resp, "ID must number")
		fmt.Printf("[ArtistHandler.getByIdHandler] Error when %w", err)
		return
	}

	artist, err := a.artistUsecase.ArtistByGenre(id)
	if err != nil {
		response.HandleError(resp, "ID must number")
		fmt.Printf("[ArtistHandler.getByIdHandler] Error when %w", err)
		return
	}
	//response.HandleSuccess(resp, artist)
	jsonData, err := json.Marshal(artist)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[ShipHandler.handleSuccess] Error when do json marshalling for error handling: %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
	return
}

func (a *ArtistHandler) insertHandler(resp http.ResponseWriter, req *http.Request) {
	reqbody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[ArtistHandler.insertHandler] Error when reading body: " + err.Error())
		return
	}

	var artist = model.Artist{}
	err = json.Unmarshal(reqbody, &artist)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[ArtistHandler.insertHandler] Error when reading body: " + err.Error())
		return
	}

	err = a.artistUsecase.Insert(&artist)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[ArtistHandler.insertHandler] Error when reading body: " + err.Error())
		return
	}
	response.HandleSuccess(resp, nil)
}

func (a *ArtistHandler) updateHandler(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[ArtistHandler.updateHandler] Error when convert: " + err.Error())
		return
	}

	data := model.Artist{}
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[ArtistHandler.updateHandler] Error when reading: " + err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[ArtistHandler.updateHandler] Error when unmarshall: " + err.Error())
		return
	}

	artist := a.artistUsecase.Update(id, &data)
	response.HandleSuccess(resp, artist)
}

func (a *ArtistHandler) deleteHandler(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[ArtistHandler.updateHandler] Error when: " + err.Error())
		return
	}

	artist := a.artistUsecase.Delete(id)
	response.HandleSuccess(resp, artist)
}

func CreateArtistHandler(r *mux.Router, artistUsecase artist.ArtistUsecase) {

	artistHandler := ArtistHandler{artistUsecase}

	r.HandleFunc("/artist", artistHandler.getAllHandler).Methods(http.MethodGet)
	s := r.PathPrefix("/artist").Subrouter()
	s.HandleFunc("/{id}", artistHandler.getByIdHandler).Methods(http.MethodGet)
	s.HandleFunc("/genre/{id}", artistHandler.getArtistByGenreHandler).Methods(http.MethodGet)
	s.HandleFunc("/add", artistHandler.insertHandler).Methods(http.MethodPost)
	s.HandleFunc("/{id}", artistHandler.updateHandler).Methods(http.MethodPut)
	s.HandleFunc("/{id}", artistHandler.deleteHandler).Methods(http.MethodDelete)
}

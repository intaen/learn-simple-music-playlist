package handler

import (
	"fmt"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/gorilla/mux"

	"spookify/genre"
	"spookify/model"
	"spookify/response"

)

type GenreHandler struct {
	genreUsecase genre.GenreUsecase
}

func (g *GenreHandler) getAllHandler(resp http.ResponseWriter, req *http.Request) {

	genre, err := g.genreUsecase.GetAll()
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Printf("[GenreHandler.getAllHandler] Error when: %w", err)
		return
	}
	//response.HandleSuccess(resp, genre)
	jsonData, err := json.Marshal(genre)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[ShipHandler.handleSuccess] Error when do json marshalling for error handling: %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
	return
}

func (g *GenreHandler) getByIdHandler(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		response.HandleError(resp, "ID must number")
		fmt.Printf("[GenreHandler.getByIdHandler] Error when %w", err)
		return
	}

	genre, err := g.genreUsecase.GetByID(id)
	if err != nil {
		response.HandleError(resp, "ID must number")
		fmt.Printf("[GenreHandler.getByIdHandler] Error when %w", err)
		return
	}
	//response.HandleSuccess(resp, artist)
	jsonData, err := json.Marshal(genre)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[ShipHandler.handleSuccess] Error when do json marshalling for error handling: %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
	return
}

func (g *GenreHandler) insertHandler(resp http.ResponseWriter, req *http.Request) {
	reqbody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[GenreHandler.insertHandler] Error when reading body: " + err.Error())
		return
	}

	var genre = model.Genre{}
	err = json.Unmarshal(reqbody, &genre)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[GenreHandler.insertHandler] Error when reading body: " + err.Error())
		return
	}

	err = g.genreUsecase.Insert(&genre)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[GenreHandler.insertHandler] Error when reading body: " + err.Error())
		return
	}
	response.HandleSuccess(resp, nil)
}

func (g *GenreHandler) updateHandler(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[GenreHandler.updateHandler] Error when convert: " + err.Error())
		return
	}

	data := model.Genre{}
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[GenreHandler.updateHandler] Error when reading: " + err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[GenreHandler.updateHandler] Error when unmarshall: " + err.Error())
		return
	}

	genre := g.genreUsecase.Update(id, &data)
	response.HandleSuccess(resp, genre)
}

func (g *GenreHandler) deleteHandler(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		response.HandleError(resp, "Sorry, something went wrong")
		fmt.Println("[GenreHandler.updateHandler] Error when: " + err.Error())
		return
	}

	genre := g.genreUsecase.Delete(id)
	response.HandleSuccess(resp, genre)
}

func CreateGenreHandler(r *mux.Router, genreUsecase genre.GenreUsecase) {

	genreHandler := GenreHandler{genreUsecase}

	r.HandleFunc("/genre", genreHandler.getAllHandler).Methods(http.MethodGet)
	s := r.PathPrefix("/genre").Subrouter()
	s.HandleFunc("/{id}", genreHandler.getByIdHandler).Methods(http.MethodGet)
	s.HandleFunc("/add", genreHandler.insertHandler).Methods(http.MethodPost)
	s.HandleFunc("/{id}", genreHandler.updateHandler).Methods(http.MethodPut)
	s.HandleFunc("/{id}", genreHandler.deleteHandler).Methods(http.MethodDelete)
}
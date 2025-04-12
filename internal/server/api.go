package server

import (
	"encoding/json"
	"fmt"
	"log"
	"microservice1/pkg/models"
	"microservice1/pkg/models/postgre"
	"net/http"
)

type API struct {
	LogInfo   *log.Logger
	LogError  *log.Logger
	Magazines *postgre.MagazineModel
}

func NewAPI() *API {
	return &API{
		LogInfo:   initLoggerInfo(),
		LogError:  initLoggerError(),
		Magazines: &postgre.MagazineModel{DB: initDb()},
	}
}

func (a *API) routes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("GET /getMagazinesByCity/{city}", a.getMagazinesByCity)
	r.HandleFunc("POST /addMagazine", a.addMagazine)
	return r
}

func (a *API) getMagazinesByCity(w http.ResponseWriter, r *http.Request) {
	city := r.PathValue("city")
	magazines, err := a.Magazines.Get(city)
	if err != nil {
		a.LogError.Printf("func getMagazinesByCity, a.Magazines.Get %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(magazines); err != nil {
		a.LogError.Printf("func getMagazinesByCity, json.Encode %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (a *API) addMagazine(w http.ResponseWriter, r *http.Request) {
	var magazine models.Magazine
	if err := json.NewDecoder(r.Body).Decode(&magazine); err != nil {
		a.LogError.Printf("func addMagazine, decode %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	id, err := a.Magazines.Insert(magazine.Name, magazine.City)
	if err != nil {
		a.LogError.Printf("func addMagazine, a.Magazines.Insert %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("Magazine %s in city %s added to DB by id %d", magazine.Name, magazine.City, id)

	a.LogInfo.Println(message)
	JSONmessage := map[string]any{"message": message, "id": id}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(JSONmessage); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}

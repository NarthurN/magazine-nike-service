package server

import (
	"log"
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
	r.HandleFunc("/getMagazinesByCity", a.getMagazinesByCity)
	r.HandleFunc("/addMagazine", a.addMagazine)
	return r
}

func (a *API) getMagazinesByCity(w http.ResponseWriter, r *http.Request) {

}

func (a *API) addMagazine(w http.ResponseWriter, r *http.Request) {
	id, err := a.Magazines.Insert("Nike", "Moscow")
	if err != nil {
		a.LogError.Println(err)
		return
	}

	a.LogInfo.Println("Nike", "Moscow", "id=", id)
}

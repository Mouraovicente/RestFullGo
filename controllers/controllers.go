package controllers

import (
	"RestFullGo/databate"
	"RestFullGo/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")

}

func GetAllPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	databate.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPersonalidadesById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade
	databate.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)

}

func CreatePersonalidade(w http.ResponseWriter, r *http.Request) {

	var p models.Personalidade
	json.NewDecoder(r.Body).Decode(&p)
	databate.DB.Create(&p)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)

}

func DeletePersonalidade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade
	databate.DB.Delete(&p, id)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(p)
}

func EditPersonalidade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade
	databate.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	databate.DB.Save(&p)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(p)
}

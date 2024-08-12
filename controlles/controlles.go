package controlles

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "github.com/Led-New/Go_apit_Rert/data-base"
	"github.com/Led-New/Go_apit_Rert/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home pagina")
}
func TodasAsPersonalidade(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}
func RertonarPersonalidae(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var persona models.Personalidade
	database.DB.First(&persona, id)
	json.NewEncoder(w).Encode(persona)
}
func CriandoPersonalidae(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)
}
func DeletaDb(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidades models.Personalidade
	database.DB.Delete(&personalidades, id)
	json.NewEncoder(w).Encode(personalidades)
}
func EditaDb(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidades models.Personalidade
	database.DB.First(&personalidades, id)
	json.NewDecoder(r.Body).Decode(&personalidades)
	database.DB.Save(&personalidades)
	json.NewEncoder(w).Encode(personalidades)
}

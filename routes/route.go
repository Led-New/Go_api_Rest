package routes

import (
	"log"
	"net/http"

	"github.com/Led-New/Go_apit_Rert/controlles"
	"github.com/Led-New/Go_apit_Rert/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//"github.com/gorilla/handlers"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.Use( middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controlles.Home)
	r.HandleFunc("/api/personalidades", controlles.TodasAsPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controlles.RertonarPersonalidae).Methods("Get")
	r.HandleFunc("/api/personalidades", controlles.CriandoPersonalidae).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controlles.DeletaDb).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controlles.EditaDb).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
package routes

import (
	"RestFullGo/controllers"
	"RestFullGo/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.GetAllPersonalidades).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controllers.GetPersonalidadesById).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CreatePersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/delete/{id}", controllers.DeletePersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/editar/{id}", controllers.EditPersonalidade).Methods("Put")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ricardomaricato/api-go-rest/controllers"
	"github.com/ricardomaricato/api-go-rest/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaUmaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}

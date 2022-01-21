package routes

import (
	"log"
	"net/http"

	"github.com/ricardomaricato/api-go-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

package main

import (
	"fmt"

	"github.com/ricardomaricato/api-go-rest/database"
	"github.com/ricardomaricato/api-go-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}

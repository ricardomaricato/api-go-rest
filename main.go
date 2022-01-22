package main

import (
	"fmt"

	"github.com/ricardomaricato/api-go-rest/database"
	"github.com/ricardomaricato/api-go-rest/models"
	"github.com/ricardomaricato/api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "nome 2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Resto com Go")
	routes.HandleRequest()
}

package main

import (
	"fmt"

	"github.com/ricardomaricato/api-go-rest/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Resto com Go")
	routes.HandleRequest()
}

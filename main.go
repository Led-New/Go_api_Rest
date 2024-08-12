package main

import (
	"fmt"

	"github.com/Led-New/Go_apit_Rert/data-base"
	"github.com/Led-New/Go_apit_Rert/models"
	"github.com/Led-New/Go_apit_Rert/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{ID: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{ID: 2, Nome: "Nome 1", Historia: "Historia 1"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor")
	routes.HandleResquest()
}

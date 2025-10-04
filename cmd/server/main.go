package main

import (
	"log"
	"oracle-crud-api/internal/config"
	db2 "oracle-crud-api/internal/db"
	"oracle-crud-api/internal/handler"
	repository "oracle-crud-api/internal/repo"
	"oracle-crud-api/internal/router"
	"oracle-crud-api/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	conn, err := db2.ConnectToOracleDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	personRepo := repository.NewPersonRepository(conn)
	personService := service.NewPersonService(personRepo)
	personHandler := handler.NewPersonHandler(personService)

	r := router.SetupRouter(personHandler)

	r.Run(":8080")
}

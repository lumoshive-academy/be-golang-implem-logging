package main

import (
	"fmt"
	"log"
	"net/http"
	"session-19/database"
	"session-19/handler"
	"session-19/repository"
	"session-19/router"
	"session-19/service"
	"session-19/utils"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	logger, err := utils.InitLogger("./logs/app-", true)

	repo := repository.NewRepository(db, logger)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	r := router.NewRouter(handler, service, logger)

	fmt.Println("server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("error server")
	}
}

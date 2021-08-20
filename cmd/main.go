package main

import (
	hotelManager "github.com/edamiyan/hotel-manager"
	"github.com/edamiyan/hotel-manager/pkg/handler"
	"github.com/edamiyan/hotel-manager/pkg/repository"
	"github.com/edamiyan/hotel-manager/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(hotelManager.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

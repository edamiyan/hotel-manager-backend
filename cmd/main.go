package main

import (
	hotelManager "github.com/edamiyan/hotel-manager"
	"github.com/edamiyan/hotel-manager/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(hotelManager.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

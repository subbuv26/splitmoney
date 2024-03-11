package main

import (
	"log"
	authController "splitmoney/internal/controllers/auth"
	authService "splitmoney/internal/services/auth"
	"splitmoney/pkg/repos/db/inmem"
)

func main() {
	userRepo := inmem.NewUserStore()
	authSvc := authService.New(userRepo)
	authApp, err := authController.NewApp(authSvc)
	if err != nil {
		log.Fatal(err)
	}
	authApp.Start()
}

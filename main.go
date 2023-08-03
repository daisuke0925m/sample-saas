package main

import (
	"log"

	"github.com/daisuke0925m/sample-saas/infrastructure"
	"github.com/daisuke0925m/sample-saas/interfaces/controllers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	userController := controllers.NewUserController(infrastructure.NewSqlHandler())
	r := infrastructure.NewRouter(userController)
	r.Run()
}

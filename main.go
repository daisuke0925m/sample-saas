package main

import (
	"github.com/daisuke0925m/sample-saas/infrastructure"
	"github.com/daisuke0925m/sample-saas/interfaces/controllers"
)

func main() {
	userController := controllers.NewUserController(infrastructure.NewSqlHandler())
	r := infrastructure.NewRouter(userController)
	r.Run()
}

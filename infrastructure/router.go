package infrastructure

import (
	"github.com/daisuke0925m/sample-saas/interfaces/controllers"
	gin "github.com/gin-gonic/gin"
)

func NewRouter(controller *controllers.UserController) *gin.Engine {
	router := gin.Default()

	router.POST("/users", func(c *gin.Context) { controller.Create(c) })
	router.GET("/users", func(c *gin.Context) { controller.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { controller.Show(c) })

	return router
}

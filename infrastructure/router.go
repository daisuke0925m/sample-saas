package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/Anti-Pattern-Inc/saasus-sdk-go/callback"
	"github.com/Anti-Pattern-Inc/saasus-sdk-go/middleware"
	"github.com/daisuke0925m/sample-saas/interfaces/controllers"
	gin "github.com/gin-gonic/gin"
)

type IDTokenGetter struct{}

func (g IDTokenGetter) GetIDToken(r *http.Request) string {
	// リクエストパラメータから　idToken を取得する
	return r.FormValue("idToken")
}

func NewRouter(controller *controllers.UserController) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.AuthMiddlewareGin(IDTokenGetter{}))

	router.GET("/callback", func(c *gin.Context) {
		fmt.Println("callback")
		callback.CallbackRouteFunction(c, c.Writer, c.Request)
	})
	// saasusのuserinfoを返す
	router.GET("/me", func(c *gin.Context) {
		c.JSON(200, c.Value("userInfo"))
	})

	router.POST("/users", func(c *gin.Context) { controller.Create(c) })
	router.GET("/users", func(c *gin.Context) { controller.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { controller.Show(c) })

	return router
}

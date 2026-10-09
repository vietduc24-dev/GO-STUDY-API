package initialize

import (
	c "go-ecommerce-backend-api/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/pong", c.NewPongController().Pong)
		v1.GET("/user/:id", c.NewUserController().GetUserByID)
	}
	return r
}

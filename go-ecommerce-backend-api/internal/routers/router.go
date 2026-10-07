package routers

import (
	c "go-ecommerce-backend-api/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/pong", c.NewPongController().Pong)
		v1.GET("/user/:id", c.NewUserController().GetUserByID)
	}
	return r
}

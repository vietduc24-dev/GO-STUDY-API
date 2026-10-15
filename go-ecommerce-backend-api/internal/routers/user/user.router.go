package user

import (
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	//this is non-denpendency
	userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp", func(c *gin.Context) { c.Status(501) })
		userRouterPublic.POST("/verify-otp", func(c *gin.Context) { c.Status(501) })
		userRouterPublic.POST("/login", func(c *gin.Context) { c.Status(501) })

	}
	//private router
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info", func(c *gin.Context) { c.Status(501) })
	}
}

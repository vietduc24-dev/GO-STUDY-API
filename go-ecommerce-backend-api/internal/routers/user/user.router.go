package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")
		userRouterPublic.POST("/verify-otp")
		userRouterPublic.POST("/login")

	}
	//private router
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
		userRouterPrivate.POST("/otp")
		userRouterPrivate.POST("/verify-otp")
		userRouterPrivate.POST("/login")

	}
}

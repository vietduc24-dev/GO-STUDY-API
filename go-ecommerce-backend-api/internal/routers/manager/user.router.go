package manager

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	//private router
	userRouterPrivate := Router.Group("/admin/user")
	//adminRouterPrivate.Use(limiter())
	//adminRouterPrivate.Use(Authen())
	//adminRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/active_user", func(c *gin.Context) { c.Status(501) })
	}
}

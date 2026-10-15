package manager

import (
	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {

	adminRouterPublic := Router.Group("/admin/user")
	{
		adminRouterPublic.POST("/login", func(c *gin.Context) { c.Status(501) })
	}
	//private router
	_ = Router.Group("/admin/user")
	//adminRouterPrivate.Use(limiter())
	//adminRouterPrivate.Use(Authen())
	//adminRouterPrivate.Use(Permission())
	{
		// /active_user is registered by UserRouter.
	}
}

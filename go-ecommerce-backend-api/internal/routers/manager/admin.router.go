package manager

import (
	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {

	adminRouterPublic := Router.Group("/admin/user")
	{
		adminRouterPublic.POST("/login")
	}
	//private router
	adminRouterPrivate := Router.Group("/admin/user")
	//adminRouterPrivate.Use(limiter())
	//adminRouterPrivate.Use(Authen())
	//adminRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/active_user")
	}
}

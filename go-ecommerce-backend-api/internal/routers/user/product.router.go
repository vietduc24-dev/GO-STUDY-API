package user

import (
	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/search", func(c *gin.Context) { c.Status(501) })
		productRouterPublic.GET("/detail/:id", func(c *gin.Context) { c.Status(501) })

	}
	//private router
}

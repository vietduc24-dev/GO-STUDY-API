package initialize

import (
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	//middleware
	// r.Use() //loggin
	// r.Use() //cross
	// r.Use() // limiter global
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("v1/2024")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)

	}
	{
		managerRouter.InitAdminRouter(MainGroup)
		managerRouter.InitUserRouter(MainGroup)
	}
	return r
}

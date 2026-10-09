package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"

	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("mysql host: ", global.Config.MySQL)
	InitLogger()
	global.Logger.Info("config log ok", zap.String("ok ", "ok"))
	InitMySql()
	InitRedis()
	r := InitRouter()

	r.Run(":8002")
}

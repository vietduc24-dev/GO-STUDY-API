package global

import (
	"go-ecommerce-backend-api/pkg/logger"
	"go-ecommerce-backend-api/pkg/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
)

/*
redis
datababase
payment...
*/

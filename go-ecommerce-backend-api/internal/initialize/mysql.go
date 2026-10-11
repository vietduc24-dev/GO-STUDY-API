package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errorString string) {
	if err != nil {
		global.Logger.Error(errorString, zap.Any("error", err))
		panic(err)
	}
}
func InitMySql() {
	m := global.Config.MySQL
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.User, m.Password, m.Host, m.Port, m.DbName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "InitMysql failed")
	global.Mdb = db

	//set pool
	SetPool()
	MigrateTables()
}

// open pool speed up performance query db
func SetPool() {
	m := global.Config.MySQL
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error: %s::", err)
	}
	//when request thi pool connection se run speedup query
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))

}

func MigrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Printf("Migrate faild")
	}
}

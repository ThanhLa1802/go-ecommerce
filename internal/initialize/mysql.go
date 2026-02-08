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

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMySQL() {
	m := global.Config.Mysql

	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	checkErrorPanic(err, "InitMysql initialization error")

	global.Logger.Info("Initializing MySQL Successfully")
	global.Mdb = db

	// set Pool
	SetPool()
	migrateTables()
}

func SetPool() {
	m := global.Config.Mysql

	sqlDB, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error: %v\n", err)
		return
	}

	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Println("Migration table error!")
	}
}

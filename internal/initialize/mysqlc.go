package initialize

import (
	"database/sql"
	"fmt"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/po"
	"time"

	"go.uber.org/zap"
)

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMySQLC() {
	m := global.Config.Mysql

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Dbname,
	)

	db, err := sql.Open("mysql", dsn)
	checkErrorPanicC(err, "Init MySQL error")

	// verify connection
	err = db.Ping()
	checkErrorPanicC(err, "Cannot connect to MySQL")

	// set pool config
	db.SetMaxIdleConns(m.MaxIdleConns)
	db.SetMaxOpenConns(m.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)

	global.Mdbc = db

	global.Logger.Info("MySQL connected successfully")

	migrateTablesC()
}

func SetPoolC() {
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

func migrateTablesC() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Println("Migration table error!")
	}
}

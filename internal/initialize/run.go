package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading config MySQL", global.Config.Mysql.Username)
	// InitLogger()
	// InitMySql()
	// InitRedis()

	r := InitRouter()

	r.Run(":8002")
}

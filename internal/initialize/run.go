package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"

	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading config MySQL", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config logger successfully!", zap.String("ok", "success"))
	InitMySQL()
	InitRedis()
	InitKafka()

	r := InitRouter()

	r.Run(":8002")
}

package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go-ecommerce-backend-api/global"
)

func LoadConfig() {
	v := viper.New()
	v.AddConfigPath("./configs")
	v.SetConfigName("local")
	v.SetConfigType("yaml")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("cannot read config: %w", err))
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable load configuration %s", err)
	}

}

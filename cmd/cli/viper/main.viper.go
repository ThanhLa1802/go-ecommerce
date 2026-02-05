package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DbName   string `mapstructure:"dbName"`
	} `mapstructure:"databases"`
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath("./configs")
	v.SetConfigName("local")
	v.SetConfigType("yaml")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil

}

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	fmt.Printf("Server Port: %d\n", cfg.GetInt("server.port"))
	var config Config

	if err := cfg.Unmarshal(&config); err != nil {
		fmt.Printf("Unable load configuration %s", err)
	}
	fmt.Printf("Config Port: %d", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("Database user: %s, Password: %s, Host: %s", db.User, db.Password, db.DbName)
	}
}

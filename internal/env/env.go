package env

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv            string `mapstructure:"APP_ENV"`
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout    int    `mapstructure:"CONTEXT_TIMEOUT"`
	Host              string `mapstructure:"HOST"`
	Port              string `mapstructure:"PORT"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPass            string `mapstructure:"DB_PASS"`
	DBName            string `mapstructure:"DB_NAME"`
	FeedFetchInterval int    `mapstructure:"INTERVAL_SECONDS"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	err := viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}

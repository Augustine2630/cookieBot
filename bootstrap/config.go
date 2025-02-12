package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	App struct {
		Token            string
		SpoonacularToken string
	} `mapstructure:"app"`
}

func NewConfig() *Config {

	//Сделано для локал разработки
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	viper.AutomaticEnv()

	// Set the prefix for environment variables (optional)
	viper.SetEnvPrefix("ABOBA")

	// Bind the environment variable to the configuration key
	viper.BindEnv("app.token", "ABOBA_TOKEN")
	viper.BindEnv("app.spoonacularToken", "ABOBA_SPOONACULAR_TOKEN")

	// Unmarshal the configuration into the struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &config
}

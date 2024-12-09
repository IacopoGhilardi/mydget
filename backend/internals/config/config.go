package config

import "github.com/spf13/viper"

type Config struct {
	DBUser      string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
	BackendPort string
}

var config *Config

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config = &Config{
		DBUser:      viper.GetString("DB_USER"),
		DBPassword:  viper.GetString("DB_PASSWORD"),
		DBHost:      viper.GetString("DB_HOST"),
		DBPort:      viper.GetString("DB_PORT"),
		DBName:      viper.GetString("DB_NAME"),
		BackendPort: viper.GetString("BACKEND_PORT"),
	}

	return config, nil
}

func GetConfig() *Config {
	return config
}

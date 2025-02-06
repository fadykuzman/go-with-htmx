package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./internal/config")
	viper.AutomaticEnv()

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading cofig: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}
	dbConfig := viper.Get("database")
	fmt.Println(dbConfig)
	fmt.Println(config.Database)

	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &config, nil
}

func validateConfig(config *Config) error {
	if config.Database.Host == "" || config.Database.User == "" ||
		config.Database.Password == "" || config.Database.Name == "" {
		return fmt.Errorf("database configuration incomplete")
	}
	return nil
}

func setDefaults() {
	db := DefaultDatabaseConfig()
	viper.SetDefault("database", db)
}

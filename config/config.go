package config

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Log *zap.Logger
	Db  *sql.DB
}

func newConfig(log *zap.Logger, db *sql.DB) *Config {
	return &Config{
		Log: log,
		Db:  db,
	}
}

var logger *zap.Logger

func Init() (*Config, error) {
	// Initialising logger
	err := initializeLog()
	if err != nil {
		fmt.Println("error", err.Error())
		return nil, err
	}

	// initialising viper configuration file
	err = initialiseViper()
	if err != nil {
		logger.Error("Error initialising viper", zap.Any("viper", err))
		return nil, err
	}

	db, err := initDB()
	if err != nil {
		logger.Error("Error initialising db", zap.Any("DB", err))
		return nil, err
	}

	return newConfig(logger, db), nil
}

func initialiseViper() error {
	viper.SetConfigFile("./conf/.env")
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Error reading .env file: %w", err)
	}
	return nil
}

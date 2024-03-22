package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func InitConfig() (*AppConfig, error) {
	ErrMissingConfig := errors.New("missing config")

	// load env
	log.Println("Trying to load .env file...")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, using values from environment variables.")
	}

	appPort, isExist := os.LookupEnv("APP_PORT")
	if !isExist {
		return nil, ErrMissingConfig
	}
	_, err = strconv.Atoi(appPort)
	if err != nil {
		return nil, errors.New("APP_PORT must be a number")
	}

	dbHost, isExist := os.LookupEnv("DB_HOST")
	if !isExist {
		return nil, ErrMissingConfig
	}

	dbPort, isExist := os.LookupEnv("DB_PORT")
	if !isExist {
		return nil, ErrMissingConfig
	}
	_, err = strconv.Atoi(dbPort)
	if err != nil {
		return nil, errors.New("DB_PORT must be a number")
	}

	dbUser, isExist := os.LookupEnv("DB_USER")
	if !isExist {
		return nil, ErrMissingConfig
	}

	dbPassword, isExist := os.LookupEnv("DB_PASSWORD")
	if !isExist {
		return nil, ErrMissingConfig
	}

	dbName, isExist := os.LookupEnv("DB_NAME")
	if !isExist {
		return nil, ErrMissingConfig
	}

	jwtSecret, isExist := os.LookupEnv("JWT_SECRET")
	if !isExist {
		return nil, ErrMissingConfig
	}

	return &AppConfig{
		AppPort:    appPort,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		JWTSecret:  jwtSecret,
	}, nil
}

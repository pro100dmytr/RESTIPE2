package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	todo "webRESTIPE2"
	"webRESTIPE2/pkg/handler"
	"webRESTIPE2/pkg/repository"
	"webRESTIPE2/pkg/service"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("init config error: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.DB_HOST"),
		Port:     viper.GetString("db.DB_PORT"),
		Username: viper.GetString("db.DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.DB_NAME"),
		SSLMode:  viper.GetString("db.DB_SSLMODE"),
	})
	if err != nil {
		logrus.Fatalf("init db error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while starting server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

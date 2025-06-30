package main

import (
	"context"
	"github.com/Numbone/golang-todo-app"
	"github.com/Numbone/golang-todo-app/pkg/handler"
	"github.com/Numbone/golang-todo-app/pkg/repository"
	"github.com/Numbone/golang-todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatal("init config err:", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal("load .env file err:", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatal("init db err:", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(golang_todo_app.Server)
	go func() {
		if err := srv.Run(":"+viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Println("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logrus.Println("server stopped")

	if err = srv.Stop(context.Background()); err != nil {
		logrus.Fatalf("error occured while stopping server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured while closing db connection: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

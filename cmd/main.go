package main

import (
	"L0"
	"L0/pkg/handler"
	"L0/pkg/repository"
	"L0/pkg/service"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	c := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	db, err := repository.NewPostgresDB(c)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	httpService := service.NewHTTPService(repos)
	natsService := service.NewNatsService(repos, viper.GetString("nats.clusterID"), viper.GetString("nats.clientID"))
	handlers := handler.NewHandler(httpService)

	err = handlers.OrderService.PullOrders()
	if err != nil {
		logrus.Fatalf("failed to pull orders: %s", err.Error())
	}

	srv := L0.NewServer(viper.GetString("port"), handlers, natsService)

	if err = srv.Run(); err != nil {
		logrus.Fatalf("error occured running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

package main

import (
	"github.com/iZillu/cryptoWallet"
	"github.com/iZillu/cryptoWallet/pkg/handler"
	"github.com/iZillu/cryptoWallet/pkg/repository"
	"github.com/iZillu/cryptoWallet/pkg/service"
	"github.com/joho/godotenv"
	mw "github.com/labstack/echo/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

func main() {
	// setting logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2 Jan 15:04:05"}).With().Caller().Logger()

	if err := initConfig(); err != nil {
		log.Fatal().Err(err).Msg("config reading")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("loading env variables")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("initializing db")
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	myEcho := handlers.InitRoutes()

	// apply logger
	//myEcho.Use(mw.LoggerWithConfig(mw.LoggerConfig{
	//	Format:           "[${time_custom}] ${status} ${method} ${uri} (${remote_ip}) ${err} ${latency_human}\n",
	//	Output:           myEcho.Logger.Output(),
	//	CustomTimeFormat: "2 Jan 15:04:05",
	//}))
	serv := new(cryptoWallet.Server)
	//myEcho.Use(handlers.UserIdentity)

	if err := serv.Start(viper.GetString("port"), myEcho); err != nil {
		log.Fatal().Err(err).Msg("cryptoWallet starting")
	}
	myEcho.Use(mw.Recover())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

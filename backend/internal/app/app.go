package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/internal/handlers"
	"residential-registration/backend/internal/services"
	"residential-registration/backend/internal/storages"
	"residential-registration/backend/pkg/database"
	"residential-registration/backend/pkg/logging"
	"syscall"
	"time"

	gormLogger "gorm.io/gorm/logger"
)

func Run() {
	conf, err := config.Init()
	if err != nil {
		log.Fatal("failed to init config", "errs", err)
	}

	logger := logging.New(conf.App.LogLevel, conf.App.SaveLogsToFile)
	logger.Info("config", "config", conf)

	defaultLocation, err := time.LoadLocation(conf.TimeLocation)
	if err != nil {
		logger.Fatal("failed to init default time location", "err", err)
	}
	time.Local = defaultLocation

	sqlLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormLogger.Error,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	sql, err := database.NewPostgreSQL(database.PostgreSQLConfig{
		User:     conf.PostgreSQL.User,
		Password: conf.PostgreSQL.Password,
		Host:     conf.PostgreSQL.Host,
		Port:     conf.PostgreSQL.Port,
		Database: conf.PostgreSQL.Database,
	}, database.SetLogger(sqlLogger))
	if err != nil {
		return
	}

	err = sql.DB.AutoMigrate(
		&entity.OSBB{},
		&entity.User{},
		&entity.Token{},
		&entity.Building{},
		&entity.Apartment{},
		&entity.Announcement{},
		&entity.Poll{},
		&entity.TestAnswer{},
		&entity.Answer{},
		&entity.Payment{},
		&entity.Purchase{},
	)
	if err != nil {
		return
	}

	storage := services.Storages{
		User:      storages.NewUserStorage(sql.DB),
		Building:  storages.NewBuildingStorage(sql.DB),
		Apartment: storages.NewApartmentStorage(sql.DB),
		OSBB:      storages.NewOSBBStorage(sql.DB),
		Token:     storages.NewTokenStorage(sql.DB),
	}
	serviceOptions := services.Options{
		Logger:   logger,
		Config:   conf,
		Storages: storage,
	}
	service := services.Services{
		Auth:  services.NewAuthService(&serviceOptions),
		Token: services.NewTokenService(&serviceOptions),
		OSBB:  services.NewOSBBService(&serviceOptions),
	}
	handler := handlers.Handler{
		Logger:   logger.Named("Handler"),
		Services: service,
	}

	srv := new(Server)
	go func() {
		if err := srv.Run(conf.Server.Port, handler.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}

}

package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/shandysiswandi/gokit-service/endpoint"
	"github.com/shandysiswandi/gokit-service/pkg/env"
	"github.com/shandysiswandi/gokit-service/pkg/logger"
	"github.com/shandysiswandi/gokit-service/repository/postgres"
	"github.com/shandysiswandi/gokit-service/service"
	"github.com/shandysiswandi/gokit-service/transport/httptrans"
)

func main() {
	// setup environment variable
	if err := godotenv.Load(".env"); err != nil {
		logger.Info("godotenv.Load " + err.Error())
		os.Exit(1)
	}

	// setup timezone
	os.Setenv("TZ", env.Get("APP_TZ"))

	// setup logger
	logger.SetOutput(os.Stderr)

	// setup repository databse
	dbRW, err := postgres.NewPostgres(postgres.Configuration{
		Host:     env.Get("DB_HOST"),
		Port:     env.GetInt("DB_PORT"),
		Username: env.Get("DB_USERNAME"),
		Password: env.Get("DB_PASSWORD"),
		Database: env.Get("DB_DATABASE"),
		Options:  env.Get("DB_OPTIONS"),
	})
	if err != nil {
		logger.Info("postgres.NewPostgres " + err.Error())
		os.Exit(1)
	}

	// setup repository cache
	// var cacheRW repository.CacheReaderWriter

	// setup business logic
	srv := service.NewTodoService(dbRW)

	// setup endpoints (controller)
	end := endpoint.NewEndpoints(srv)

	// setup tranport (router)
	h := httptrans.NewServer(end, env.Get("JWT_SECRET"))

	// setup http server
	server := &http.Server{
		Handler:      h,
		Addr:         "0.0.0.0:" + env.Get("APP_PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 30,
	}

	// running server in groutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Info("server.ListenAndServe " + err.Error())
			os.Exit(1)
		}
	}()
	logger.Info("server listen on port " + env.Get("APP_PORT"))

	// setup channel for listen some trigger
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// get channel value, it will blocking until channel is writing some value
	<-c

	// setup context with timeout for shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
		dbRW.Close()
	}()

	// shutdown server
	if err := server.Shutdown(ctx); err != nil {
		logger.Info("server.Shutdown " + err.Error())
		os.Exit(1)
	}

	logger.Info("shutting down server")
}

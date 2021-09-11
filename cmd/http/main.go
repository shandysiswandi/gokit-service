package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joho/godotenv"
	"github.com/shandysiswandi/gokit-service/endpoint"
	"github.com/shandysiswandi/gokit-service/pkg/env"
	"github.com/shandysiswandi/gokit-service/repository"
	"github.com/shandysiswandi/gokit-service/repository/postgres"
	"github.com/shandysiswandi/gokit-service/repository/redis"
	"github.com/shandysiswandi/gokit-service/service"
	"github.com/shandysiswandi/gokit-service/transport/httptrans"
)

func main() {
	// setup environment variable
	if err := godotenv.Load(".env"); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	// setup timezone
	os.Setenv("TZ", env.Get("APP_TZ"))

	// setup logger
	var logger log.Logger
	{
		logger = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
		timezone := log.DefaultTimestampUTC
		if env.Get("APP_TZ") != "UTC" {
			timezone = log.DefaultTimestamp
		}
		logger = log.With(logger, "ts", timezone)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// setup repository databse
	var dbRW repository.DatabaseReaderWriter
	var err error
	{
		dbRW, err = postgres.NewPostgres(postgres.Configuration{
			Host:     env.Get("DB_HOST"),
			Port:     env.GetInt("DB_PORT"),
			Username: env.Get("DB_USERNAME"),
			Password: env.Get("DB_PASSWORD"),
			Database: env.Get("DB_DATABASE"),
			Options:  env.Get("DB_OPTIONS"),
		})
		if err != nil {
			level.Error(logger).Log("msg", err)
			os.Exit(1)
		}
	}

	// setup repository cache
	var cacheRW repository.CacheReaderWriter
	{
		cacheRW, err = redis.NewRedis()
		if err != nil {
			level.Error(logger).Log("msg", err)
			os.Exit(1)
		}
	}

	// setup business logic
	var srv service.TodoService
	{
		srv = service.NewTodoService(logger, dbRW, cacheRW)
		srv = service.NewMiddleware(logger)(srv)
	}

	// setup endpoints (controller)
	var end endpoint.Endpoints
	{
		end = endpoint.NewEndpoints(srv, env.Get("JWT_SECRET"))
	}

	// setup tranport (router)
	var h http.Handler
	{
		h = httptrans.NewServer(end)
	}

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
			level.Error(logger).Log("msg", err)
			os.Exit(1)
		}
	}()
	level.Info(logger).Log("msg", "server listen on port "+env.Get("APP_PORT"))

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
		level.Error(logger).Log("msg", err)
		os.Exit(1)
	}

	level.Info(logger).Log("msg", "shutting down server ")
}

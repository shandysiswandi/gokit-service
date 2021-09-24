package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joho/godotenv"
	"github.com/shandysiswandi/gokit-service/endpoint"
	"github.com/shandysiswandi/gokit-service/pkg/env"
	v1 "github.com/shandysiswandi/gokit-service/proto"
	"github.com/shandysiswandi/gokit-service/repository"
	"github.com/shandysiswandi/gokit-service/repository/postgres"
	"github.com/shandysiswandi/gokit-service/repository/redis"
	"github.com/shandysiswandi/gokit-service/service"
	"github.com/shandysiswandi/gokit-service/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// setup environment variable
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
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
		cacheRW, err = redis.NewRedis(redis.Configuration{
			Host:     env.Get("CACHE_HOST"),
			Port:     env.GetInt("CACHE_PORT"),
			Password: env.Get("CACHE_PASSWORD"),
			DB:       env.GetInt("CACHE_DB"),
		})
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

	// Create net listener
	grpcListener, err := net.Listen("tcp", "0.0.0.0:"+env.Get("APP_PORT"))
	if err != nil {
		level.Error(logger).Log("msg", err)
		os.Exit(1)
	}

	// Create grpc options
	opts, err := grpcOptions()
	if err != nil {
		level.Error(logger).Log("msg", err)
		os.Exit(1)
	}

	// serve grpc server in goroutine
	go func() {
		grpcServer := transport.NewServer(end)
		server := grpc.NewServer(opts...)

		v1.RegisterTodoServiceServer(server, grpcServer)

		level.Info(logger).Log("msg", "Server listen on port "+env.Get("APP_PORT"))
		server.Serve(grpcListener)
	}()

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("exit %s", <-c)
	}()

	level.Error(logger).Log("msg", <-errs)
}

func grpcOptions() ([]grpc.ServerOption, error) {
	opts := []grpc.ServerOption{}

	// Create the TLS credentials
	cert, err := tls.LoadX509KeyPair(env.Get("CERT_PATH"), env.Get("KEY_PATH"))
	if err != nil {
		return nil, err
	}

	if env.Get("APP_ENV") != "production" {
		return opts, nil
	}

	opts = []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	return opts, nil
}

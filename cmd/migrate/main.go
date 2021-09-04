package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"github.com/shandysiswandi/gokit-service/pkg/env"
	"github.com/shandysiswandi/gokit-service/repository/postgres"
)

func main() {
	// setup environment variable
	if err := godotenv.Load(".env"); err != nil {
		log.Println("err", err)
		os.Exit(1)
	}

	log.Println("⏩ running migration...")
	dbRW, err := postgres.NewPostgres(postgres.Configuration{
		Host:     env.Get("DB_HOST"),
		Port:     env.GetInt("DB_PORT"),
		Username: env.Get("DB_USERNAME"),
		Password: env.Get("DB_PASSWORD"),
		Database: env.Get("DB_DATABASE"),
		Options:  env.Get("DB_OPTIONS"),
	})
	if err != nil {
		log.Println("err", err)
		os.Exit(1)
	}

	// auto migration
	if err := goose.Up(dbRW.GetDB(), "resource/migration"); err != nil {
		log.Println("err", err)
		os.Exit(1)
	}

	log.Println("💯 migration done.")
}

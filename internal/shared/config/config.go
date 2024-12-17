package config

import (
	"AnonymusBot/internal/shared/storage/dto"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func GetSlogConfig() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func GetRedis() *dto.RedisConfig {
	addr := os.Getenv("REDIS_ADDR")
	pass := os.Getenv("REDIS_PASSWORD")
	if addr == "" || pass == "" {
		log.Fatal("Не удалось получить данные для подключения к Redis из файла .env")
	}

	return &dto.RedisConfig{
		Addr:     addr,
		Password: pass,
	}
}

func GetPostgres() (*dto.PostgresConfig, error) {
	addr := os.Getenv("POSTGRES_CONNECT_STRING")
	if addr == "" {
		slog.Error("Не удалось получить данные для подключения к Postgres из файла .env ")
		return nil, fmt.Errorf("не удалось получить данные для подключения к Postgres из файла .env ")
	}

	return &dto.PostgresConfig{Addr: addr}, nil
}

func GetHTTPAddress() *dto.AddressHTTP {
	addr := os.Getenv("HTTP_ADDRESS")
	if addr == "" {
		addr = ":8086"
	}

	return &dto.AddressHTTP{Addr: addr}
}

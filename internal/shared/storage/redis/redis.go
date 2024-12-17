package redis

import (
	"AnonymusBot/internal/shared/config"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"log/slog"
)

func InitRedis(ctx context.Context) *redis.Client {
	cfg := config.GetRedis()

	rDB := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       0,
	})

	if _, err := rDB.Ping(ctx).Result(); err != nil {
		log.Fatalf("Не удалось подключиться к Redis: %v", err)
	}

	slog.Info("Подключение к Redis успешно.")

	return rDB
}

package postgres

import (
	"AnonymusBot/internal/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

func InitPostgres() (*gorm.DB, error) {
	cfg, err := config.GetPostgres()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(cfg.Addr), &gorm.Config{})
	if err != nil {
		slog.Error("Не удалось подключиться к Postgres.", "Ошибка", err)
		return nil, err
	}

	if err := db.AutoMigrate(); err != nil {
		slog.Error("Не удалось подключиться к Postgres.", "Ошибка", err)
		return nil, err
	}

	return db, nil
}

package postgres

import (
	"customer-service/config"
	"customer-service/internal/domain/models"
	"customer-service/pkg/logger"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDb(config *config.Config) (*gorm.DB, error) {
	const op = "storage.postgres.New"

	var dbUrl string

	if config.Mode == "Local" {
		dbUrl = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
			config.Postgres.PostgresHost, config.Postgres.PostgresPort,
			config.Postgres.PostgresUser, config.Postgres.PostgresPassword,
			config.Postgres.PostgresDbname)
	} else {
		dbUrl = os.Getenv("DATABASE_URL")
	}

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if config.Mode != "prod" {
		db = db.Debug()
	}

	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		logger.Logger.Fatal().Str(op, err.Error()).Msg("Pg init failed")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.AutoMigrate(&models.CustomerModel{}); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}

package database

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/config"
	"go.uber.org/zap"
	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func Connect(cfg *config.DBConfig) *gorm.DB {
	zap.L().Debug("Database connection initializing.")

	db, err := gorm.Open(psql.Open(cfg.DataSourceName), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("Cannot connect to database", zap.Error(err))
	}

	origin, err := db.DB()
	if err != nil {
		zap.L().Fatal("Cannot get sql.DB from database", zap.Error(err))
	}

	origin.SetMaxOpenConns(cfg.MaxOpen)
	origin.SetMaxIdleConns(cfg.MaxIdle)
	origin.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)

	zap.L().Debug("Database connection successfully initialized.")
	return db
}

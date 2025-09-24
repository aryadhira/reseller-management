package database

import (
	"fmt"
	"time"

	"github.com/aryadhira/reseller-management/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewDB(cfg *DBConfig) (*gorm.DB, error) {
	strConn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port, "disable",
	)

	db, err := gorm.Open(postgres.Open(strConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// auto Migration
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	pg, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Connection pool settings
	pg.SetMaxOpenConns(25)
	pg.SetMaxIdleConns(25)
	pg.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

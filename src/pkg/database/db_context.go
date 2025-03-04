package database

import (
	"fmt"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log/slog"
	"time"
)

type Options func(ctx *dbContext)
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

var DbContextSet = wire.NewSet(NewDbContext)

type dbContext struct {
	config *PostgresConfig
}

func NewDbContext(config *PostgresConfig) DbContext {
	return &dbContext{
		config: config,
	}
}

func (ctx *dbContext) GetDatabase() (*gorm.DB, error) {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", ctx.config.Host, ctx.config.Username, ctx.config.Password, ctx.config.Database, ctx.config.Port)
	slog.Info(conn)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return db, nil
}

func (ctx *dbContext) Configure(options ...Options) {
	for _, option := range options {
		option(ctx)
	}
}

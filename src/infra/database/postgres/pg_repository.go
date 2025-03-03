package postgres

import (
	"fmt"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"go-foodshop/src/infra/database"
	"go-foodshop/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var RepositorySet = wire.NewSet(NewPgRepository[models.Food])

type (
	PostgresConfig struct {
		Host     string `env-required:"true" yaml:"host"    env:"HOST"`
		Port     int    `env-required:"true" yaml:"port"    env:"PORT"`
		Username string `env-required:"true" yaml:"username" env:"USERNAME"`
		Password string `env-required:"true" yaml:"password" env:"PASSWORD"`
		Database string `env-required:"true" yaml:"database" env:"DATABASE"`
	}
)

type pgRepository[T any] struct {
	db *gorm.DB
}

func (pg *pgRepository[T]) AddEntity(context echo.Context, entity *T) (*T, error) {
	pg.db.Create(entity)
	return entity, nil
}

func (pg *pgRepository[T]) RemoveEntity(context echo.Context, id int) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func (pg *pgRepository[T]) GetAllEntities(context echo.Context) ([]*T, error) {
	data := make([]*T, 0)
	pg.db.Find(&data)
	return data, nil
}

func (pg *pgRepository[T]) GetEntityById(context echo.Context, id int) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func NewPgRepository[T any](config *PostgresConfig) (database.Repository[T], error) {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.Username, config.Password, config.Database, config.Port)
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

	db.AutoMigrate(&models.Food{}, &models.Category{})
	return &pgRepository[T]{
		db: db,
	}, nil
}

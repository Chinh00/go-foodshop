package database

import (
	"gorm.io/gorm"
)

type Database interface {
	GetDatabase() (*gorm.DB, error)
}

package database

import "gorm.io/gorm"

type DbContext interface {
	GetDatabase() (*gorm.DB, error)
	Configure(...Options)
}

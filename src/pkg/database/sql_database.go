package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ Database = (*SqlDatabase)(nil)

type SqlDatabase struct {
	dbConnection *gorm.DB
}

func (sqlDatabase *SqlDatabase) New(conn string) Database {
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to database ...")
	sqlDatabase.dbConnection = db
	return sqlDatabase
}

func (sqlDatabase *SqlDatabase) GetDatabase() (*gorm.DB, error) {
	return sqlDatabase.dbConnection, nil
}

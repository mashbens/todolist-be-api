package util

import (
	"fmt"

	"github.com/mashbens/todolist/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseDriver string

const (
	MYSQL DatabaseDriver = "MYSQL"
)

type DatabaseConnection struct {
	Driver DatabaseDriver

	MYSQL *gorm.DB
}

func NewConnectionDatabase(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection

	switch config.Driver {
	case "MYSQL":
		db.Driver = MYSQL
		db.MYSQL = NewMYSQL(config)
	default:
		panic("Database driver not supported")
	}
	return &db
}
func NewMYSQL(config *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_User,
		config.DB_Pass,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Debug().Msg(dsn)
	return db
}

func (db *DatabaseConnection) CloseConnection() {
	if db.MYSQL != nil {
		db, _ := db.MYSQL.DB()
		db.Close()
	}
}

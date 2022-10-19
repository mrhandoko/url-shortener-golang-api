package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB Config
type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

func Connect(config Config) (dbConn *gorm.DB, err error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.DB)
	dbConn, err = gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	log.Println("DB Connected")
	return dbConn, nil
}

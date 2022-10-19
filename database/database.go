package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Connector *gorm.DB

// DB Config
type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

var GetConnectionString = func(config Config) string {
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.DB)
	return connString
}

func Connect(connString string) (err error) {
	Connector, err = gorm.Open("mysql", connString)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	log.Println("DB Connected")
	return nil
}

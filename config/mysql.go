package config

import (
	"fmt"
	"os"
	"ticketing/model/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func ConnectDB() *gorm.DB {
	username := os.Getenv("APP_DB_USERNAME")
	if username == "" {
		username = "root"
	}
	password := os.Getenv("APP_DB_PASSWORD")
	if password == "" {
		password = ""
	}
	port := os.Getenv("APP_DB_PORT")
	if port == "" {
		port = "3306"
	}
	host := os.Getenv("APP_DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	name := os.Getenv("APP_DB_NAME")
	if name == "" {
		name = "alta"
	}

	config := Config{
		DB_Username: username,
		DB_Password: password,
		DB_Port:     port,
		DB_Host:     host,
		DB_Name:     name,
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitialMigration(DB)
	return DB
}

func InitialMigration(DB *gorm.DB) {
	DB.AutoMigrate(
		&domain.User{},
		&domain.Ticket{},
		&domain.Order{},
		&domain.Event{},
		&domain.OrderDetail{},
	)
}

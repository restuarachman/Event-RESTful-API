package config

import (
	"fmt"
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
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"",
		"localhost",
		"3306",
		"alta",
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

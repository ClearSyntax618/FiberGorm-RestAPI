package config

import (
	"github.com/ClearSyntax618/FiberGorm-RestAPI/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DB_URI string = "root:root@tcp(localhost:3001)/dogs-mysql?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
	var err error

	DB, err = gorm.Open(mysql.Open(DB_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Dog{})

	return nil
}

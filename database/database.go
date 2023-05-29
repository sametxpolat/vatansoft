package database

import (
	"fmt"

	"github.com/sametxpolat/vatansoft/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionMySQL() *gorm.DB {
	dsn := "neo:01382nN*@tcp(127.0.0.1:3306)/vatansoft@localhost?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("> server-msg: successful to connect database")

	err = db.AutoMigrate(&model.Product{}, &model.Category{})
	if err != nil {
		panic(err)
	}

	return db
}

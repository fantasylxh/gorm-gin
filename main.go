package main

import (
	"gorm-gin/Config"
	"gorm-gin/Models"
	"gorm-gin/Routers"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", "express:nkds3EZPtYzJP2mt@tcp(127.0.0.1:3306)/sdrms?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Book{})
	Config.DB.LogMode(true)
	r := Routers.SetupRouter()
	// running
	r.Run(":8082")
}

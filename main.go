package main

import (
<<<<<<< HEAD
	"./Config"
	"./Models"
	"./Routers"
=======
	"gorm-gin/Config"
	"gorm-gin/Models"
	"gorm-gin/Routers"
>>>>>>> b9011c4f04cea88061361ff0a17d69a3c16690c3
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

<<<<<<< HEAD
	Config.DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/testinger?charset=utf8&parseTime=True&loc=Local")
=======
	Config.DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sdrms?charset=utf8&parseTime=True&loc=Local")
>>>>>>> b9011c4f04cea88061361ff0a17d69a3c16690c3

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Book{})
<<<<<<< HEAD

	r := Routers.SetupRouter()
	// running
	r.Run()
=======
	Config.DB.LogMode(true)
	r := Routers.SetupRouter()
	// running
	r.Run(":8082")
>>>>>>> b9011c4f04cea88061361ff0a17d69a3c16690c3
}

package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm-gin/Config"
)

func UpdateUser(b *BackendUser, id string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func GetUser(b *BackendUser, user_name string, user_pwd string) (err error) {
	if err := Config.DB.Where("user_name = ? and user_pwd = ?", user_name, user_pwd).First(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUser(b *BackendUser, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

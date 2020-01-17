package Models

import (
	"gorm-gin/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllOrder(b *[]Order) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewOrder(b *Order) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneOrder(b *Order, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneOrder(b *Order, id string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteOrder(b *Order, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}

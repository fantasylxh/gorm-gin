package Models

import (
	"gorm-gin/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllAddress(b *[]Address) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewAddress(b *Address) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneAddress(b *Address, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneAddress(b *Address, id string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteAddress(b *Address, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}

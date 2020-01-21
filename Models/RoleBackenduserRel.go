package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm-gin/Config"
)

//根据uid获取角色id 24:快递管理员-孟加拉 25:快递管理员-中国
func GetOneRole(b *BackendUser, id string) (err error) {
	if err := Config.DB.Order("id desc").Where("backend_user_id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

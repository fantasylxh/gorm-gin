package Models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Author   string `json:"author" form:"author"`
	Category string `json:"category" form:"category"`
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryPrice struct {
	Price string `json:"price"`
}

type BackendUser struct {
	Id       int    `json:"uid" form:"uid" primaryKey:"true"`
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	RealName string `json:"real_name,omitempty"`
	//AccessToken string `json:"access_token" form:"token" binding:"required"`
	UserPwd string `json:"-" "user_pwd" form:"user_pwd"`
	Status  int    `json:"status"`
}

type Order struct {
	Id           int       `json:"order_id"` // 自增
	OrderSn      string    `json:"order_sn" form:"order_sn" binding:"required"`
	Weight       string    `json:"weight" form:"weight" binding:"required"`
	CategoryId   string    `json:"category_id" form:"category_id" binding:"required"`
	PayTime      string    `json:"pay_time"`
	OrderType    string    `json:"order_type" form:"order_type"`
	OrderStatus  int       `json:"order_status"`
	OrderPrice   string    `json:"order_price"`
	PayStatus    string    `json:"pay_status"`
	ShipStatus   string    `json:"ship_status"`
	PayQrcode    string    `json:"pay_qrcode" form:"pay_qrcode"`
	Country      string    `json:"country" form:"country"`
	Province     string    `json:"province" form:"province"`
	City         string    `json:"city" form:"city"`
	Address      string    `json:"address" form:"address"`
	Mobile       string    `json:"mobile" form:"mobile"`
	AddressId    int       `json:"address_id" form:"address_id"`
	RecAddressId int       `json:"rec_address_id" form:"rec_address_id"`
	RecCountry   string    `json:"rec_country" form:"rec_country"`
	RecProvince  string    `json:"rec_province" form:"rec_province"`
	RecCity      string    `json:"rec_city" form:"rec_city"`
	RecAddress   string    `json:"rec_address" form:"rec_address"`
	RecMobile    string    `json:"rec_mobile" form:"rec_mobile"`
	CreatedAt    time.Time `json:"created_at"`
	CreatorId    int       `json:"creator_id"`
}

func (b *Book) TableName() string {
	return "book"
}

func (b *Category) TableName() string {
	return "rms_category"
}

func (b *CategoryPrice) TableName() string {
	return "rms_category_price"
}

func (b *BackendUser) TableName() string {
	return "rms_backend_user"
}

func (b *Order) TableName() string {
	return "rms_order"
}


package Models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}

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
	Id           int          `json:"order_id"` // 自增
	OrderSn      string       `json:"order_sn" form:"order_sn" binding:"required"`
	Weight       string       `json:"weight" form:"weight" binding:"required"`
	CategoryId   string       `json:"category_id" form:"category_id" binding:"required"`
	PayTime      Time         `json:"pay_time"`
	OrderType    string       `gorm:"default:'unknown'" "column:order_type"`
	OrderStatus  int          `json:"order_status"`
	OrderPrice   string       `gorm:"default:'0.00'" "column:order_price"`
	PayStatus    int          `gorm:"default:'0'" column:"pay_status"`
	ShipStatus   int          `gorm:"default:0" column:"ship_status"`
	PayQrcode    string       `json:"pay_qrcode" form:"pay_qrcode"`
	Country      string       `json:"country" form:"country"`
	Province     string       `json:"province" form:"province"`
	City         string       `json:"city" form:"city"`
	Address      string       `json:"address" form:"address"`
	Mobile       string       `json:"mobile" form:"mobile"`
	AddressId    int          `json:"address_id" form:"address_id"`
	RecAddressId int          `json:"rec_address_id" form:"rec_address_id"`
	RecCountry   string       `json:"rec_country" form:"rec_country"`
	RecProvince  string       `json:"rec_province" form:"rec_province"`
	RecCity      string       `json:"rec_city" form:"rec_city"`
	RecAddress   string       `json:"rec_address" form:"rec_address"`
	RecMobile    string       `json:"rec_mobile" form:"rec_mobile"`
	CreatedAt    Time         `json:"created_at"`
	UpdatedAt    time.Time    `json:"-" "updated_at"`
	CreatorId    string       `json:"creator_id"`
	ShipActions  []ShipAction `gorm:"FOREIGNKEY:OrderId;ASSOCIATION_FOREIGNKEY:Id"`
}
type ShipAction struct {
	Id         int    `json:"id"` // 自增
	ShipStatus int    `json:"ship_status"`
	ActionNote string `json:"action_note"`
	CreatedAt  Time   `json:"created_at"`
	DeletedAt  Time   `json:"-" deleted_at"`
	OrderId    int    `json:"order_id"`
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
func (b *ShipAction) TableName() string {
	return "rms_ship_action"
}

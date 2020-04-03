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
	OrderCode    string       `json:"order_code" form:"order_code"`
	OrderStatus  int          `json:"order_status" form:"order_status"`
	OrderPrice   string       `json:"order_price"`
	PayStatus    int          `json:"pay_status" form:"pay_status"`
	ShipStatus   int          `json:"ship_status" form:"ship_status"`
	PayQrcode    string       `json:"pay_qrcode" form:"pay_qrcode"`
	Country      string       `json:"country" form:"country"`
	Province     string       `json:"province" form:"province"`
	City         string       `json:"city" form:"city"`
	Address      string       `json:"address" form:"address"`
	Mobile       string       `json:"mobile" form:"mobile"`
	AddressId    string       `json:"address_id" form:"address_id"`
	RecAddressId string       `json:"rec_address_id" form:"rec_address_id"`
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
	Id         int       `json:"id"` // 自增
	ShipStatus int       `json:"ship_status"`
	ActionNote string    `json:"action_note"`
	CreatedAt  string    `json:"created_at"` //time.Time时间格式化输出问题 todo 优化
	DeletedAt  time.Time `json:"-" deleted_at"`
	OrderId    int       `json:"order_id" form:"order_id" binding:"required"`
	CreatorId  int       `json:"creator_id"`
}

type OrderAction struct {
	ID         uint
	ActionNote string
	CreatedAt  time.Time
	DeletedAt  time.Time
	OrderId    int `json:"order_id" form:"order_id" binding:"required"`
	CreatorId  int
}

type Address struct {
	ID          uint   `json:"address_id"`
	Country     string `json:"country" form:"country" binding:"required"`
	Province    string `json:"province" form:"province" binding:"required"`
	City        string `json:"city" form:"city" binding:"required"`
	Address     string `json:"address" form:"address" binding:"required"`
	Mobile      string `json:"mobile" form:"mobile" binding:"required"`
	AddressType int    `json:"address_type" form:"address_type"`
	IsDefault   int    `json:"is_default" form:"is_default"`
	CreatorId   int    `json:"creator_id"`
	CreatedAt   Time   `json:"created_at"`
	UpdatedAt   Time   `json:"-" "updated_at"`
	DeletedAt   Time   `json:"-" "deleted_at"`
}
type PaymentCode struct {
	ID        uint   `json:"id"`
	Img       string `json:"img"`
	OrderCode string `json:"order_code" form:"order_code" binding:"required"`
	Status    string `json:"status" form:"status"`
	CreatedAt Time   `json:"-" "created_at"`
	UpdatedAt Time   `json:"-" "updated_at"`
	DeletedAt Time   `json:"-" "deleted_at"`
}

type FeedBack struct {
	gorm.Model
	Content   string `json:"content" form:"content" binding:"required"`
	CreatorId string
}

type RoleBackenduserRel struct {
	RoleId        string `json:"role_id"`
	BackendUserId string `json:"backend_user_id" form:"backend_user_id"`
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

func (b *OrderAction) TableName() string {
	return "rms_order_action"
}

func (b *ShipAction) TableName() string {
	return "rms_ship_action"
}

func (b *Address) TableName() string {
	return "rms_user_address"
}

func (b *PaymentCode) TableName() string {
	return "rms_payment_code"
}

func (b *FeedBack) TableName() string {
	return "rms_feedback"
}

func (b *RoleBackenduserRel) TableName() string {
	return "rms_role_backenduser_rel"
}

package Models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Author   string `json:"author" form:"author"`
	Category string `json:"category" form:"category"`
}

func (b *Book) TableName() string {
	return "book"
}

package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title  string
	UserId int
	User   User    `gorm:"foreignKey:UserId"`
}
//Color  []Color `gorm:"many2many:product_color;"`
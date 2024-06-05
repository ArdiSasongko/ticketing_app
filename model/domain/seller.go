package domain

import "time"

type Sellers struct {
	SellerID  int     `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string  `gorm:"column:name"`
	Email     string  `gorm:"column:email"`
	Password  string  `gorm:"column:password"`
	Events    []Event `gorm:"foreignKey:SellerID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
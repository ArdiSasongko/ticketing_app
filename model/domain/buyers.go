package domain

import "time"

type Buyers struct {
	BuyerID   int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

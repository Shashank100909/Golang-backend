package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

type Student struct {
	Id      int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name    string `json:"name"`
	Class   int    `json:"class"`
	Section string `json:"section"`
}

type Products struct {
	Id          int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string  `gorm:"not null;unique" json:"name" binding:"required"`
	Price       float64 `gorm:"not null" json:"price" binding:"required"`
	Description string  `json:"description,omitempty"`
}

type Cart struct {
	UserID    int       `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;references:ID"`
	ProductID int       `json:"product_id"`
	Product   Products  `gorm:"foreignKey:ProductID;references:Id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}


package models

import (
	"time"
)

type User struct {
	UserID       int    `gorm:"primaryKey;autoIncrement" json:"user_id"`
	FirstName    string `json:"first_name"  binding:"required"`
	LastName     string `json:"last_name"  binding:"required"`
	Age          int    `json:"age"  binding:"required"`
	Gender       string `json:"gender"  binding:"required"`
	Email        string `json:"email"  binding:"required"`
	MobileNumber int    `json:"mobile_number"  binding:"required"`
	Username     string `json:"username"  binding:"required"`
	Password     string `json:"password"  binding:"required"`
	// Carts        []Cart `gorm:"foreignKey:UserID;references:UserID"`
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
	User      User      `gorm:"foreignKey:UserID;references:UserID"`
	ProductID int       `json:"product_id"`
	Product   Products  `gorm:"foreignKey:ProductID;references:Id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

type Address struct {
	AddressID    int    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       int    `json:"user_id"`
	Name         string `json:"name"`
	MobileNumber int    `json:"mobile_number"`
	Street       string `json:"street"`
	Landmark     string `json:"landmark"`
	City         string `json:"city"`
	Pincode      int    `json:"pincode"`
	State        string `json:"state"`
	Country      string `json:"country"`
}

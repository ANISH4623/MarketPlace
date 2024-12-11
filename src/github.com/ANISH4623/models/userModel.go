package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"gorm_._model"`
	Username   string    `json:"username" validate:"required,unique" gorm:"unique;not null"` // Enforce uniqueness and not null
	FirstName  string    `json:"first_name" validate:"required" json:"first_name,omitempty"`
	LastName   string    `json:"last_name" validate:"required" json:"last_name,omitempty"`
	Email      string    `json:"email" validate:"required,email,unique" gorm:"unique;not null"` // Enforce uniqueness and not null
	Password   string    `json:"password" validate:"required" json:"password,omitempty"`
	Address    string    `json:"Address" json:"address,omitempty"`
	Products   []Product `json:"products" gorm:"many2many:product_products;" json:"products,omitempty"`
}

type Product struct {
	gorm.Model
	ProductID   int     `json:"product_id" validate:"required,gte=1" gorm:"primary_key"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	Quantity    int64   `json:"quantity" validate:"required"`
	Image       string  `json:"Image"`
	UserID      uint    `json:"user_id"`
}

package models

import "time"

type Product struct {
	Id          int64      `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	Name        string     `gorm:"not null; size:255; unique;" json:"name"`
	Price       float64    `gorm:"not null;" json:"price"`
	Amount      int64      `gorm:"not null;" json:"amount"`
	Description string     `gorm:"size:255;" json:"description"`
	CreatedAt   *time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"not null;" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`

	Categories []Category `json:"categories" gorm:"many2many:category_products;"`
	Images     []Image    `json:"images"`
}

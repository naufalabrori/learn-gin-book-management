package models

import "time"

type Book struct {
	ID                uint      `json:"id" gorm:"primary_key"`
	Title             string    `json:"title" gorm:"size:150"`
	Author            string    `json:"author" gorm:"size:150"`
	Publisher         string    `json:"publisher" gorm:"size:150"`
	PublishedYear     string    `json:"published_year" gorm:"size:5"`
	ISBN              string    `json:"isbn" gorm:"size:13;unique"`
	CategoryID        uint      `json:"category_id" gorm:"not null"`
	Quantity          uint      `json:"quantity" gorm:"not null"`
	AvailableQuantity uint      `json:"available_quantity" gorm:"not null"`
	CreatedDate       time.Time `json:"created_date" gorm:"autoCreateTime"`
	ModifiedDate      time.Time `json:"modified_date" gorm:"autoUpdateTime"`
}

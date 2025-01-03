package models

import "time"

type Book struct {
	ID                uint      `json:"id" gorm:"primary_key"`
	Title             string    `json:"title" gorm:"size:150"`
	Author            string    `json:"author" gorm:"size:150"`
	Publisher         string    `json:"publisher" gorm:"size:150"`
	PublishedYear     string    `json:"publishedDate" gorm:"size:5"`
	ISBN              string    `json:"isbn" gorm:"size:13"`
	CategoryID        uint      `json:"categoryId" gorm:"not null"`
	Quantity          uint      `json:"quantity" gorm:"not null"`
	AvailableQuantity uint      `json:"availableQuantity" gorm:"not null"`
	CreatedDate       time.Time `json:"createdDate" gorm:"autoCreateTime"`
	ModifiedDate      time.Time `json:"modifiedDate" gorm:"autoUpdateTime"`
}

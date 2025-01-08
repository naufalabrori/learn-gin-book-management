package models

import "time"

type Category struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	CategoryName string    `json:"category_name" gorm:"size:100;not null;unique"`
	CreatedDate  time.Time `json:"created_date" gorm:"autoCreateTime"`
	ModifiedDate time.Time `json:"modified_date" gorm:"autoUpdateTime"`
}

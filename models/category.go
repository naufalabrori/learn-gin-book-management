package models

import "time"

type Category struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	CategoryName string    `json:"categoryName" gorm:"size:100;not null;unique"`
	CreatedDate  time.Time `json:"createdDate" gorm:"autoCreateTime"`
	ModifiedDate time.Time `json:"modifiedDate" gorm:"autoUpdateTime"`
}

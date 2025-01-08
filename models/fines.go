package models

import "time"

type Fines struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	TransactionID uint      `json:"transaction_id" gorm:"not null"`
	Amount        float64   `json:"amount" gorm:"not null"`
	PaidDate      time.Time `json:"paid_date" gorm:"autoCreateTime"`
	CreatedDate   time.Time `json:"created_date" gorm:"autoCreateTime"`
	ModifiedDate  time.Time `json:"modified_date" gorm:"autoUpdateTime"`
}

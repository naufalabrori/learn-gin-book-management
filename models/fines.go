package models

import "time"

type Fines struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	TransactionID uint      `json:"transactionId" gorm:"not null"`
	Amount        float64   `json:"amount" gorm:"not null"`
	PaidDate      time.Time `json:"paidDate"`
	CreatedDate   time.Time `json:"createdDate" gorm:"autoCreateTime"`
	ModifiedDate  time.Time `json:"modifiedDate" gorm:"autoUpdateTime"`
}

package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name             string    `gorm:"type:varchar(100);not null"`
	Email            string    `gorm:"type:varchar(100);unique;not null"`
	RegistrationDate time.Time `gorm:"type:timestamp"`
	Loans            []Loan    `gorm:"foreignKey:UserID"`
}

type Users []User

package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title             string `gorm:"type:varchar(100);not null"`
	Author            string `gorm:"type:varchar(100);not null"`
	PublicationYear   int    `gorm:"type:int;not null"`
	AvailableQuantity int    `gorm:"type:int;not null"`
	Loans             []Loan `gorm:"foreignKey:BookID"`
}

type Books []Book

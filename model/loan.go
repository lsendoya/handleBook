package model

import (
	"gorm.io/gorm"
	"time"
)

type Loan struct {
	gorm.Model
	LoanDate   time.Time  `gorm:"type:timestamp;not null"`
	ReturnDate time.Time  `gorm:"type:timestamp"`
	Status     LoanStatus `gorm:"type:varchar(50);not null"`
	UserID     uint       `gorm:"not null"`
	BookID     uint       `gorm:"not null"`
	User       User       `gorm:"foreignKey:UserID"`
	Book       Book       `gorm:"foreignKey:BookID"`
}

type Loans []Loan

type LoanStatus string

const (
	Pending   LoanStatus = "Pending"
	Active    LoanStatus = "Active"
	Overdue   LoanStatus = "Overdue"
	Renewed   LoanStatus = "Renewed"
	Completed LoanStatus = "Completed"
	Cancelled LoanStatus = "Cancelled"
	Lost      LoanStatus = "Lost"
)

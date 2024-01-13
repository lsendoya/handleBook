package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Loan struct {
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	LoanDate   time.Time      `gorm:"type:timestamp;null" json:"loan_date"`
	ReturnDate time.Time      `gorm:"type:timestamp" json:"return_date,omitempty"`
	Status     LoanStatus     `gorm:"type:varchar(50);null" json:"status"`
	UserID     uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	BookID     uuid.UUID      `gorm:"type:uuid;not null" json:"book_id"`
	User       *User          `gorm:"foreignKey:UserID" json:"user"`
	Book       *Book          `gorm:"foreignKey:BookID" json:"book"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type Loans []Loan

type LoanStatus string

type LoanUUIDS struct {
	UserID uuid.UUID `json:"user_id"`
	BookID uuid.UUID `json:"book_id"`
}

const (
	Pending   LoanStatus = "Pending"
	Active    LoanStatus = "Active"
	Overdue   LoanStatus = "Overdue"
	Renewed   LoanStatus = "Renewed"
	Completed LoanStatus = "Completed"
	Cancelled LoanStatus = "Cancelled"
	Lost      LoanStatus = "Lost"
)

func (l Loan) BeforeCreate(loan *Loan) {
	loan.ID = uuid.New()
	loan.CreatedAt = time.Now()
	loan.UpdatedAt = time.Now()
	loan.LoanDate = time.Now()
	loan.ReturnDate = loan.LoanDate.Add(time.Hour * 24 * 10)
}

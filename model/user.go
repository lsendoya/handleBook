package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID               uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name             string         `gorm:"type:varchar(100);not null" json:"name"`
	Email            string         `gorm:"type:varchar(100);unique;not null" json:"email"`
	RegistrationDate time.Time      `gorm:"type:timestamp" json:"registration_date"`
	Loans            *[]Loan        `gorm:"foreignKey:UserID" json:"loans,omitempty"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type Users []User

func (u User) BeforeCreate(user *User) {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.RegistrationDate = time.Now()
}

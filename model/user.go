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
	Password         string         `gorm:"type:varchar(256);null" json:"password"`
	Role             RoleUser       `gorm:"type:varchar(100);null" json:"role"`
	RegistrationDate time.Time      `gorm:"type:timestamp" json:"registration_date"`
	Loans            []Loan         `gorm:"foreignKey:UserID" json:"loans"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type Users []User

type RoleUser string

const (
	Admin   RoleUser = "admin"
	Regular RoleUser = "regular"
)

func (u User) BeforeCreate(user *User) {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.RegistrationDate = time.Now()
}

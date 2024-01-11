package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID                uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Title             string         `gorm:"type:varchar(100);not null" json:"title"`
	Author            string         `gorm:"type:varchar(100);not null" json:"author"`
	PublicationYear   int            `gorm:"type:int;not null" json:"publication_year"`
	AvailableQuantity int            `gorm:"type:int;not null" json:"available_quantity"`
	Publisher         string         `gorm:"type:varchar(100);not null" json:"publisher"`
	Loans             *[]Loan        `gorm:"foreignKey:BookID" json:"loans,omitempty"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type Books []Book

func (b Book) BeforeCreate(book *Book) {
	book.ID = uuid.New()
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
}

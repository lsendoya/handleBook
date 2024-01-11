package book

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lsendoya/handleBook/model"
	"gorm.io/gorm"
)

type Book struct {
	db *gorm.DB
}

func New(db *gorm.DB) Book {
	return Book{db: db}
}

func (b Book) Add(book *model.Book) error {
	if err := b.db.Create(book).Error; err != nil {
		return fmt.Errorf("error creating book %w", err)
	}
	return nil
}
func (b Book) List() (model.Books, error) {
	var books model.Books
	if err := b.db.Find(&books).Error; err != nil {
		return nil, fmt.Errorf("error retrieving books %w", err)
	}
	return books, nil
}
func (b Book) Get(bookId uuid.UUID) (*model.Book, error) {
	var book model.Book
	result := b.db.Where("id = ?", bookId).First(&book)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}

	return &book, nil
}
func (b Book) Update(bookId uuid.UUID, payload interface{}) error {
	mdl, err := b.Get(bookId)
	if err != nil {
		return err
	}

	if err := b.db.Model(mdl).Updates(payload).Error; err != nil {
		return err
	}

	return nil
}
func (b Book) Delete(bookId uuid.UUID) error {
	result := b.db.Delete(&model.Book{}, "id = ?", bookId)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

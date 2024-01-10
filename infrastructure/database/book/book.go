package book

import (
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
	return nil
}
func (b Book) List() (model.Books, error) {
	return nil, nil
}
func (b Book) Get(bookId string) (*model.Book, error) {
	return &model.Book{}, nil
}
func (b Book) Update(bookId string, payload interface{}) error {
	return nil
}
func (b Book) Delete(bookId string) error {
	return nil
}

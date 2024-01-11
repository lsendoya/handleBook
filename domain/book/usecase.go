package book

import (
	"github.com/google/uuid"
	"github.com/lsendoya/handleBook/model"
)

type Book struct {
	storage Storage
}

func New(s Storage) Book {
	return Book{s}
}

func (b *Book) Add(book *model.Book) error {
	book.BeforeCreate(book)
	return b.storage.Add(book)
}
func (b *Book) List() (model.Books, error) {
	return b.storage.List()
}
func (b *Book) Get(bookId uuid.UUID) (*model.Book, error) {
	return b.storage.Get(bookId)
}
func (b *Book) Update(bookId uuid.UUID, payload interface{}) error {
	return b.storage.Update(bookId, payload)
}
func (b *Book) Delete(bookId uuid.UUID) error {
	return b.storage.Delete(bookId)
}

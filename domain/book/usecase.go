package book

import "github.com/lsendoya/handleBook/model"

type Book struct {
	storage Storage
}

func New(s Storage) Book {
	return Book{s}
}

func (b *Book) Add(book *model.Book) error {
	return nil
}
func (b *Book) List() (model.Books, error) {
	return nil, nil
}
func (b *Book) Get(bookId string) (*model.Book, error) {
	return &model.Book{}, nil
}
func (b *Book) Update(bookId string, payload interface{}) error {
	return nil
}
func (b *Book) Delete(bookId string) error {
	return nil
}

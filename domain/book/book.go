package book

import "github.com/lsendoya/handleBook/model"

type UseCase interface {
	Add(book *model.Book) error
	List() (model.Books, error)
	Get(bookId string) (*model.Book, error)
	Update(bookId string, payload interface{}) error
	Delete(bookId string) error
}

type Storage interface {
	Add(book *model.Book) error
	List() (model.Books, error)
	Get(bookId string) (*model.Book, error)
	Update(bookId string, payload interface{}) error
	Delete(bookId string) error
}

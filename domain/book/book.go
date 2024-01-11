package book

import (
	"github.com/google/uuid"
	"github.com/lsendoya/handleBook/model"
)

type UseCase interface {
	Add(book *model.Book) error
	List() (model.Books, error)
	Get(bookId uuid.UUID) (*model.Book, error)
	Update(bookId uuid.UUID, payload interface{}) error
	Delete(bookId uuid.UUID) error
}

type Storage interface {
	Add(book *model.Book) error
	List() (model.Books, error)
	Get(bookId uuid.UUID) (*model.Book, error)
	Update(bookId uuid.UUID, payload interface{}) error
	Delete(bookId uuid.UUID) error
}

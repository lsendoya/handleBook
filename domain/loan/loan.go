package loan

import (
	"github.com/google/uuid"
	"github.com/lsendoya/handleBook/model"
)

type UseCase interface {
	Register(loan model.Loan) (*model.Loan, error)
	List() (model.Loans, error)
	UpdateStatus(loadId uuid.UUID, payload interface{}) error
	Get(loadId uuid.UUID) (*model.Loan, error)
}

type Storage interface {
	Register(loan model.Loan) (*model.Loan, error)
	List() (model.Loans, error)
	UpdateStatus(loadId uuid.UUID, payload interface{}) error
	Get(loadId uuid.UUID) (*model.Loan, error)
}

type UseCaseUser interface {
	Get(userId uuid.UUID) (*model.User, error)
}

type UseCaseBook interface {
	Get(bookId uuid.UUID) (*model.Book, error)
}

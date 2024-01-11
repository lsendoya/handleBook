package loan

import (
	"github.com/google/uuid"
	"github.com/lsendoya/handleBook/model"
)

type UseCase interface {
	Register(loan interface{}) error
	List() (model.Loans, error)
	Update(loadId uuid.UUID, payload *model.Loan) error
}

type Storage interface {
	Register(loan interface{}) error
	List() (model.Loans, error)
	Update(loadId uuid.UUID, payload *model.Loan) error
}

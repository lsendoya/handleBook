package loan

import (
	"github.com/google/uuid"
	"github.com/lsendoya/handleBook/model"
)

type Loan struct {
	storage Storage
}

func New(s Storage) Loan {
	return Loan{s}
}

func (l *Loan) Register(loan interface{}) error {
	return l.Register(loan)
}
func (l *Loan) List() (model.Loans, error) {
	return l.List()
}
func (l *Loan) Update(loadId uuid.UUID, payload *model.Loan) error {
	return l.Update(loadId, payload)
}

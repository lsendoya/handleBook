package loan

import (
	"github.com/lsendoya/handleBook/model"
	"gorm.io/gorm"
)

type Loan struct {
	db *gorm.DB
}

func New(db *gorm.DB) Loan {
	return Loan{db: db}
}

func (l Loan) Register(loan *model.Loan) error {
	return nil
}
func (l Loan) List() (model.Loans, error) {
	return nil, nil
}
func (l Loan) Update(loadId string) error {
	return nil
}

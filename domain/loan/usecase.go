package loan

import "github.com/lsendoya/handleBook/model"

type Loan struct {
	storage Storage
}

func New(s Storage) Loan {
	return Loan{s}
}

func (l *Loan) Register(loan *model.Loan) error {
	return nil
}
func (l *Loan) List() (model.Loans, error) {
	return nil, nil
}
func (l *Loan) Update(loadId string) error {
	return nil
}

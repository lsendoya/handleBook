package loan

import "github.com/lsendoya/handleBook/model"

type UseCase interface {
	Register(loan *model.Loan) error
	List() (model.Loans, error)
	Update(loadId string) error
}

type Storage interface {
	Register(loan *model.Loan) error
	List() (model.Loans, error)
	Update(loadId string) error
}

package loan

import (
	"github.com/google/uuid"
	"github.com/lsendoya/handleBook/model"
)

type Loan struct {
	storage Storage
	book    UseCaseBook
	user    UseCaseUser
}

func New(s Storage, b UseCaseBook, u UseCaseUser) Loan {
	return Loan{
		storage: s,
		book:    b,
		user:    u,
	}
}

func (l *Loan) Register(loan model.Loan) (*model.Loan, error) {

	loan.BeforeCreate(&loan)

	uuidBook := loan.BookID
	uuidUser := loan.UserID

	_, errUser := l.user.Get(uuidUser)
	if errUser != nil {
		return nil, errUser
	}

	_, errBook := l.book.Get(uuidBook)
	if errBook != nil {
		return nil, errBook
	}

	loan.Status = model.Active

	return l.storage.Register(loan)
}
func (l *Loan) List() (model.Loans, error) {
	return l.storage.List()
}
func (l *Loan) UpdateStatus(loadId uuid.UUID, payload interface{}) error {
	return l.storage.UpdateStatus(loadId, payload)
}

func (l *Loan) Get(loadId uuid.UUID) (*model.Loan, error) {
	return l.storage.Get(loadId)
}

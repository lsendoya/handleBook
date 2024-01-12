package loan

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lsendoya/handleBook/model"
	"gorm.io/gorm"
)

type Loan struct {
	db *gorm.DB
}

func New(db *gorm.DB) Loan {
	return Loan{db: db}
}

func (l Loan) Register(loan model.Loan) (*model.Loan, error) {
	if err := l.db.Create(loan).Error; err != nil {
		return nil, fmt.Errorf("error creating loan %w", err)
	}
	return nil, nil
}
func (l Loan) List() (model.Loans, error) {
	var loans model.Loans
	if err := l.db.Find(&loans).Error; err != nil {
		return nil, fmt.Errorf("error retrieving loans %w", err)
	}

	return loans, nil
}
func (l Loan) UpdateStatus(loadId uuid.UUID, payload interface{}) error {
	mdl, err := l.Get(loadId)
	if err != nil {
		return err
	}

	if err := l.db.Model(mdl).Updates(payload).Error; err != nil {
		return err
	}

	return nil
}

func (l Loan) Get(loadId uuid.UUID) (*model.Loan, error) {
	var loan model.Loan

	result := l.db.Where("id = ?", loadId).First(&loan)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}

	return &loan, nil
}

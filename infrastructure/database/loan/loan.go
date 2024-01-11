package loan

import (
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

func (l Loan) Register(loan interface{}) error {
	if err := l.db.Create(loan).Error; err != nil {
		return fmt.Errorf("error creating loan %w", err)
	}
	return nil
}
func (l Loan) List() (model.Loans, error) {
	var loans model.Loans
	if err := l.db.Find(&loans).Error; err != nil {
		return nil, fmt.Errorf("error retrieving loans %w", err)
	}

	return loans, nil
}
func (l Loan) Update(loadId uuid.UUID, payload *model.Loan) error {
	if err := l.db.Model(&model.Loan{}).Where("id = ?", loadId).Updates(payload).Error; err != nil {
		return fmt.Errorf("error updating loan: %w", err)
	}

	return nil
}

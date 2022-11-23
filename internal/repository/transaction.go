package repository

import (
	"github.com/flaiers/flallet/internal/entity"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) Transaction(
	fn func(*gorm.DB) error, tx ...*gorm.DB,
) error {
	return r.txDefault(tx...).Transaction(fn)
}

func (r *TransactionRepository) txDefault(tx ...*gorm.DB) *gorm.DB {
	if len(tx) < 1 {
		return r.db
	}

	return tx[0]
}

func (r *TransactionRepository) Create(
	transaction *entity.Transaction, tx ...*gorm.DB,
) (*entity.Transaction, error) {
	err := r.txDefault(tx...).Create(transaction).Error
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *TransactionRepository) Report(
	month uint, year uint,
) ([]*entity.Report, error) {
	var report []*entity.Report
	err := r.db.Model(&entity.Transaction{}).Select("service_id, sum(amount) as total").Where(
		"EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ? AND status = ?",
		month, year, "approved",
	).Where("service_id IS NOT NULL").Group("service_id").Find(&report).Error
	if err != nil {
		return nil, err
	}

	return report, nil
}

func (r *TransactionRepository) Find(
	conditions *entity.Transaction, scopes ...func(*gorm.DB) *gorm.DB,
) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction
	err := r.db.Scopes(scopes...).Where(
		conditions,
	).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepository) FindOneAndUpdate(
	conditions *entity.Transaction,
	values *entity.Transaction,
	tx ...*gorm.DB,
) (*entity.Transaction, error) {
	transaction := &entity.Transaction{}
	err := r.txDefault(tx...).Where(
		conditions,
	).Updates(values).First(transaction).Error
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

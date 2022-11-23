package service

import (
	"github.com/flaiers/flallet/internal/dto"
	"github.com/flaiers/flallet/internal/entity"
	"github.com/flaiers/flallet/internal/repository"
	"github.com/flaiers/flallet/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionService struct {
	repository *repository.TransactionRepository
}

func NewTransactionService(db *gorm.DB) *TransactionService {
	return &TransactionService{
		repository: repository.NewTransactionRepository(db),
	}
}

func (s *TransactionService) Create(
	transactionCreate *dto.TransactionCreate, tx ...*gorm.DB,
) (*entity.Transaction, error) {
	return s.repository.Create(
		&entity.Transaction{
			Type:      transactionCreate.Type,
			Status:    transactionCreate.Status,
			Amount:    transactionCreate.Amount,
			UserID:    transactionCreate.UserID,
			OrderID:   transactionCreate.OrderID,
			ServiceID: transactionCreate.ServiceID,
		}, tx...,
	)
}

func (s *TransactionService) Report(
	month uint, year uint,
) ([]*entity.Report, error) {
	return s.repository.Report(month, year)
}

func (s *TransactionService) FindByUserID(
	userID uuid.UUID,
	pagination *pkg.Pagination[dto.TransactionRead],
	transactionOrder *dto.TransactionOrder,
) ([]*entity.Transaction, error) {
	orderQuery := func(db *gorm.DB) *gorm.DB {
		if transactionOrder.OrderByAmount != "" {
			db = db.Order("amount " + transactionOrder.OrderByAmount)
		}
		if transactionOrder.OrderByCreatedAt != "" {
			db = db.Order("created_at " + transactionOrder.OrderByCreatedAt)
		}

		return db
	}

	return s.repository.Find(
		&entity.Transaction{UserID: userID}, pagination.Query, orderQuery,
	)
}

func (s *TransactionService) Reject(
	transactionFilter *dto.TransactionFilter, tx ...*gorm.DB,
) (*entity.Transaction, error) {
	return s.repository.FindOneAndUpdate(
		&entity.Transaction{
			Amount:    transactionFilter.Amount,
			UserID:    transactionFilter.UserID,
			OrderID:   &transactionFilter.OrderID,
			ServiceID: &transactionFilter.ServiceID,
		}, &entity.Transaction{Status: "rejected"}, tx...,
	)
}

func (s *TransactionService) Approve(
	transactionFilter *dto.TransactionFilter, tx ...*gorm.DB,
) (*entity.Transaction, error) {
	return s.repository.FindOneAndUpdate(
		&entity.Transaction{
			Amount:    transactionFilter.Amount,
			UserID:    transactionFilter.UserID,
			OrderID:   &transactionFilter.OrderID,
			ServiceID: &transactionFilter.ServiceID,
		}, &entity.Transaction{Status: "approved"}, tx...,
	)
}

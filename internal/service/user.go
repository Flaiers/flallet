package service

import (
	"github.com/flaiers/flallet/internal/dto"
	"github.com/flaiers/flallet/internal/entity"
	"github.com/flaiers/flallet/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	repository         *repository.UserRepository
	transactionService *TransactionService
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		repository:         repository.NewUserRepository(db),
		transactionService: NewTransactionService(db),
	}
}

func (s *UserService) FindOneOrCreate(
	userID uuid.UUID, tx ...*gorm.DB,
) (*entity.User, error) {
	return s.repository.FindOneOrCreate(userID, tx...)
}

func (s *UserService) AccrueUserBalance(
	userID uuid.UUID, userAccrual *dto.UserAccrual, tx ...*gorm.DB,
) (*entity.User, error) {
	var user *entity.User
	err := s.repository.Transaction(func(tx *gorm.DB) (err error) {
		user, err = s.FindOneOrCreate(userID, tx)
		if err != nil {
			return
		}
		user, err = s.repository.AccrueUserBalance(
			user, userAccrual.Amount, tx,
		)
		if err != nil {
			return
		}
		_, err = s.transactionService.Create(&dto.TransactionCreate{
			Type:      "accrual",
			Status:    "approved",
			Amount:    userAccrual.Amount,
			UserID:    userID,
			OrderID:   userAccrual.OrderID,
			ServiceID: userAccrual.ServiceID,
		}, tx)
		if err != nil {
			return
		}

		return
	}, tx...)

	return user, err
}

func (s *UserService) WithdrawUserBalance(
	userID uuid.UUID,
	userWithdrawal *dto.UserWithdrawal,
	status string,
	tx ...*gorm.DB,
) (*entity.User, error) {
	var user *entity.User
	err := s.repository.Transaction(func(tx *gorm.DB) (err error) {
		user, err = s.FindOneOrCreate(userID, tx)
		if err != nil {
			return
		}
		user, err = s.repository.WithdrawUserBalance(
			user, userWithdrawal.Amount, tx,
		)
		if err != nil {
			return
		}
		_, err = s.transactionService.Create(&dto.TransactionCreate{
			Type:      "withdrawal",
			Status:    status,
			Amount:    userWithdrawal.Amount,
			UserID:    userID,
			OrderID:   userWithdrawal.OrderID,
			ServiceID: userWithdrawal.ServiceID,
		}, tx)
		if err != nil {
			return
		}

		return
	}, tx...)

	return user, err
}

func (s *UserService) TransferUserBalance(
	userID uuid.UUID, userTransfer *dto.UserTransfer, tx ...*gorm.DB,
) (*entity.User, error) {
	var user *entity.User
	err := s.repository.Transaction(func(tx *gorm.DB) (err error) {
		user, err = s.FindOneOrCreate(userID, tx)
		if err != nil {
			return
		}
		user, err = s.repository.WithdrawUserBalance(
			user, userTransfer.Amount, tx,
		)
		if err != nil {
			return
		}
		_, err = s.transactionService.Create(&dto.TransactionCreate{
			Type:   "transfer",
			Status: "approved",
			Amount: userTransfer.Amount,
			UserID: userID,
		}, tx)
		if err != nil {
			return
		}
		_, err = s.AccrueUserBalance(
			userTransfer.UserID, &dto.UserAccrual{
				Amount: userTransfer.Amount,
			}, tx,
		)
		if err != nil {
			return
		}

		return
	}, tx...)

	return user, err
}

func (s *UserService) RefundUserBalance(
	userID uuid.UUID, userRefund *dto.UserRefund, tx ...*gorm.DB,
) (*entity.User, error) {
	var user *entity.User
	err := s.repository.Transaction(func(tx *gorm.DB) (err error) {
		user, err = s.FindOneOrCreate(userID, tx)
		if err != nil {
			return
		}
		user, err = s.repository.AccrueUserBalance(
			user, userRefund.Amount, tx,
		)
		if err != nil {
			return
		}

		return
	}, tx...)

	return user, err
}

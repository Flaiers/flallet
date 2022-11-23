package service

import (
	"github.com/flaiers/flallet/internal/dto"
	"github.com/flaiers/flallet/internal/entity"
	"github.com/flaiers/flallet/internal/repository"
	"gorm.io/gorm"
)

type ReservationService struct {
	repository         *repository.ReservationRepository
	userService        *UserService
	transactionService *TransactionService
}

func NewReservationService(db *gorm.DB) *ReservationService {
	return &ReservationService{
		repository:         repository.NewReservationRepository(db),
		userService:        NewUserService(db),
		transactionService: NewTransactionService(db),
	}
}

func (s *ReservationService) Create(
	reservationCreate *dto.ReservationCreate, tx ...*gorm.DB,
) (*entity.Reservation, error) {
	var reservation *entity.Reservation
	err := s.repository.Transaction(func(tx *gorm.DB) (err error) {
		reservation, err = s.repository.Create(
			&entity.Reservation{
				Amount:    reservationCreate.Amount,
				UserID:    reservationCreate.UserID,
				OrderID:   reservationCreate.OrderID,
				ServiceID: reservationCreate.ServiceID,
			}, tx,
		)
		if err != nil {
			return
		}
		_, err = s.userService.WithdrawUserBalance(
			reservationCreate.UserID, &dto.UserWithdrawal{
				Amount:    reservationCreate.Amount,
				OrderID:   &reservationCreate.OrderID,
				ServiceID: &reservationCreate.ServiceID,
			}, "pending", tx,
		)
		if err != nil {
			return
		}

		return
	}, tx...)

	return reservation, err
}

func (s *ReservationService) Reject(
	reservationFilter *dto.ReservationFilter, tx ...*gorm.DB,
) (*entity.Reservation, error) {
	var reservation *entity.Reservation
	err := s.repository.Transaction(func(tx *gorm.DB) (err error) {
		reservation, err = s.repository.FindOneAndDelete(
			&entity.Reservation{
				Amount:    reservationFilter.Amount,
				UserID:    reservationFilter.UserID,
				OrderID:   reservationFilter.OrderID,
				ServiceID: reservationFilter.ServiceID,
			}, tx,
		)
		if err != nil {
			return
		}
		_, err = s.userService.RefundUserBalance(
			reservationFilter.UserID, &dto.UserRefund{
				Amount: reservationFilter.Amount,
			}, tx,
		)
		if err != nil {
			return
		}
		_, err = s.transactionService.Reject(
			&dto.TransactionFilter{
				UserID:    reservationFilter.UserID,
				ServiceID: reservationFilter.ServiceID,
				OrderID:   reservationFilter.OrderID,
				Amount:    reservationFilter.Amount,
			}, tx,
		)
		if err != nil {
			return
		}

		return
	}, tx...)

	return reservation, err
}

func (s *ReservationService) Approve(
	reservationFilter *dto.ReservationFilter, tx ...*gorm.DB,
) (*entity.Reservation, error) {
	var reservation *entity.Reservation
	err := s.repository.Transaction(func(tx *gorm.DB) (err error) {
		reservation, err = s.repository.FindOneAndDelete(
			&entity.Reservation{
				Amount:    reservationFilter.Amount,
				UserID:    reservationFilter.UserID,
				OrderID:   reservationFilter.OrderID,
				ServiceID: reservationFilter.ServiceID,
			}, tx,
		)
		if err != nil {
			return
		}
		_, err = s.transactionService.Approve(
			&dto.TransactionFilter{
				UserID:    reservationFilter.UserID,
				ServiceID: reservationFilter.ServiceID,
				OrderID:   reservationFilter.OrderID,
				Amount:    reservationFilter.Amount,
			}, tx,
		)
		if err != nil {
			return
		}

		return
	}, tx...)

	return reservation, err
}

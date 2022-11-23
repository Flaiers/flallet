package repository

import (
	"github.com/flaiers/flallet/internal/entity"
	"gorm.io/gorm"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{
		db: db,
	}
}

func (r *ReservationRepository) Transaction(
	fn func(*gorm.DB) error, tx ...*gorm.DB,
) error {
	return r.txDefault(tx...).Transaction(fn)
}

func (r *ReservationRepository) txDefault(tx ...*gorm.DB) *gorm.DB {
	if len(tx) < 1 {
		return r.db
	}

	return tx[0]
}

func (r *ReservationRepository) Create(
	reservation *entity.Reservation, tx ...*gorm.DB,
) (*entity.Reservation, error) {
	err := r.txDefault(tx...).Create(reservation).Error
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (r *ReservationRepository) FindOneAndDelete(
	conditions *entity.Reservation, tx ...*gorm.DB,
) (*entity.Reservation, error) {
	reservation := &entity.Reservation{}
	err := r.txDefault(tx...).Where(
		conditions,
	).First(reservation).Delete(&entity.Reservation{}).Error
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

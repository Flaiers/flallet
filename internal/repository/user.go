package repository

import (
	"github.com/flaiers/flallet/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Transaction(
	fn func(*gorm.DB) error, tx ...*gorm.DB,
) error {
	return r.txDefault(tx...).Transaction(fn)
}

func (r *UserRepository) txDefault(tx ...*gorm.DB) *gorm.DB {
	if len(tx) < 1 {
		return r.db
	}

	return tx[0]
}

func (r *UserRepository) FindOneOrCreate(
	userID uuid.UUID, tx ...*gorm.DB,
) (*entity.User, error) {
	user := &entity.User{}
	err := r.txDefault(tx...).Where(entity.User{
		ID: userID,
	}).FirstOrCreate(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) AccrueUserBalance(
	user *entity.User, amount uint, tx ...*gorm.DB,
) (*entity.User, error) {
	user.Balance += amount
	err := r.txDefault(tx...).Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) WithdrawUserBalance(
	user *entity.User, amount uint, tx ...*gorm.DB,
) (*entity.User, error) {
	if user.Balance < amount || user.Balance == 0 {
		return nil, fiber.NewError(
			fiber.StatusPaymentRequired,
			"user balance is not enough",
		)
	}
	user.Balance -= amount
	err := r.txDefault(tx...).Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

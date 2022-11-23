package entity

import (
	"time"

	"github.com/flaiers/flallet/internal/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;unique;index;default:gen_random_uuid()"`
	Balance   uint      `gorm:"type:integer;default:0;not null"`
	CreatedAt time.Time `gorm:"default:now();not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (u *User) ToUserRead() *dto.UserRead {
	return &dto.UserRead{
		ID:        u.ID,
		Balance:   u.Balance,
		CreatedAt: u.CreatedAt,
	}
}

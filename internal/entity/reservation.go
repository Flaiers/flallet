package entity

import (
	"time"

	"github.com/flaiers/flallet/internal/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;unique;index;default:gen_random_uuid()"`
	Amount    uint      `gorm:"type:integer;not null"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null;unique"`
	ServiceID uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt time.Time `gorm:"default:now();not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (r *Reservation) ToReservationRead() *dto.ReservationRead {
	return &dto.ReservationRead{
		ID:        r.ID,
		Amount:    r.Amount,
		UserID:    r.UserID,
		OrderID:   r.OrderID,
		ServiceID: r.ServiceID,
		CreatedAt: r.CreatedAt,
	}
}

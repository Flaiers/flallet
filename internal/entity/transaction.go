package entity

import (
	"time"

	"github.com/flaiers/flallet/internal/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;unique;index;default:gen_random_uuid()"`
	Type      string     `gorm:"type:transaction_type;not null"`
	Status    string     `gorm:"type:transaction_status;not null"`
	Amount    uint       `gorm:"type:integer;not null"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null"`
	OrderID   *uuid.UUID `gorm:"type:uuid;default:null;unique"`
	ServiceID *uuid.UUID `gorm:"type:uuid;default:null"`
	CreatedAt time.Time  `gorm:"default:now();not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (t *Transaction) ToTransactionRead() *dto.TransactionRead {
	return &dto.TransactionRead{
		ID:        t.ID,
		Type:      t.Type,
		Status:    t.Status,
		Amount:    t.Amount,
		UserID:    t.UserID,
		OrderID:   t.OrderID,
		ServiceID: t.ServiceID,
		CreatedAt: t.CreatedAt,
	}
}

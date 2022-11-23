package dto

import (
	"time"

	"github.com/google/uuid"
)

type TransactionCreate struct {
	Type      string     `json:"type" enums:"accrual,withdrawal,transfer" validate:"required"`
	Status    string     `json:"status" enums:"pending,approved,rejected" validate:"required"`
	Amount    uint       `json:"amount" validate:"required,number,gt=0"`
	UserID    uuid.UUID  `json:"userId" format:"uuid" validate:"required,uuid"`
	OrderID   *uuid.UUID `json:"orderId,omitempty" format:"uuid" validate:"uuid"`
	ServiceID *uuid.UUID `json:"serviceId,omitempty" format:"uuid" validate:"uuid"`
}

type TransactionRead struct {
	ID        uuid.UUID  `json:"id" format:"uuid"`
	Type      string     `json:"type" enums:"accrual,withdrawal,transfer"`
	Status    string     `json:"status" enums:"pending,approved,rejected"`
	Amount    uint       `json:"amount"`
	UserID    uuid.UUID  `json:"userId" format:"uuid"`
	OrderID   *uuid.UUID `json:"orderId" format:"uuid"`
	ServiceID *uuid.UUID `json:"serviceId" format:"uuid"`
	CreatedAt time.Time  `json:"createdAt"`
}

type TransactionOrder struct {
	OrderByAmount    string `json:"orderByAmount" enums:"asc,desc"`
	OrderByCreatedAt string `json:"orderByCreatedAt" enums:"asc,desc"`
}

type TransactionFilter struct {
	Amount    uint      `json:"amount" validate:"required,number,gt=0"`
	UserID    uuid.UUID `json:"userId" format:"uuid" validate:"required,uuid"`
	OrderID   uuid.UUID `json:"orderId" format:"uuid" validate:"required,uuid"`
	ServiceID uuid.UUID `json:"serviceId" format:"uuid" validate:"required,uuid"`
}

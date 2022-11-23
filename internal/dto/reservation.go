package dto

import (
	"time"

	"github.com/google/uuid"
)

type ReservationCreate struct {
	Amount    uint      `json:"amount" validate:"required,number,gt=0"`
	UserID    uuid.UUID `json:"userId" format:"uuid" validate:"required,uuid"`
	OrderID   uuid.UUID `json:"orderId" format:"uuid" validate:"required,uuid"`
	ServiceID uuid.UUID `json:"serviceId" format:"uuid" validate:"required,uuid"`
}

type ReservationRead struct {
	ID        uuid.UUID `json:"id" format:"uuid"`
	Amount    uint      `json:"amount"`
	UserID    uuid.UUID `json:"userId" format:"uuid"`
	OrderID   uuid.UUID `json:"orderId" format:"uuid"`
	ServiceID uuid.UUID `json:"serviceId" format:"uuid"`
	CreatedAt time.Time `json:"createdAt"`
}

type ReservationFilter struct {
	Amount    uint      `json:"amount" validate:"required,number,gt=0"`
	UserID    uuid.UUID `json:"userId" format:"uuid" validate:"required,uuid"`
	OrderID   uuid.UUID `json:"orderId" format:"uuid" validate:"required,uuid"`
	ServiceID uuid.UUID `json:"serviceId" format:"uuid" validate:"required,uuid"`
}

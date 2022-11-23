package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserAccrual struct {
	Amount    uint       `json:"amount" validate:"required,number,gt=0"`
	OrderID   *uuid.UUID `json:"orderId,omitempty" format:"uuid" validate:"uuid"`
	ServiceID *uuid.UUID `json:"serviceId,omitempty" format:"uuid" validate:"uuid"`
}

type UserWithdrawal struct {
	Amount    uint       `json:"amount" validate:"required,number,gt=0"`
	OrderID   *uuid.UUID `json:"orderId,omitempty" format:"uuid" validate:"uuid"`
	ServiceID *uuid.UUID `json:"serviceId,omitempty" format:"uuid" validate:"uuid"`
}

type UserTransfer struct {
	Amount uint      `json:"amount" validate:"required,number,gt=0"`
	UserID uuid.UUID `json:"userId" format:"uuid" validate:"required,uuid"`
}

type UserRefund struct {
	Amount uint `json:"amount" validate:"required,number,gt=0"`
}

type UserRead struct {
	ID        uuid.UUID `json:"id" format:"uuid"`
	Balance   uint      `json:"balance"`
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
}

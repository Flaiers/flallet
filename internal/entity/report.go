package entity

import (
	"strconv"

	"github.com/google/uuid"
)

type Report struct {
	ServiceID uuid.UUID
	Total     int
}

func (r *Report) ToString() []string {
	return []string{
		r.ServiceID.String(), strconv.Itoa(r.Total),
	}
}

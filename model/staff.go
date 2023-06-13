package model

import (
	"time"

	"github.com/google/uuid"
)

type Staff struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Create_time time.Time `json:"create_time"`

	Staff_name string `json:"staff_name"`
}

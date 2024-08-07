package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type NotficationPool struct {
	gorm.Model

	UserId string `gorm:"unique"`

	TokenArray pq.StringArray `gorm:"type:text[]"`
}

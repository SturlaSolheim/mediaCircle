package models

import (
	"time"
	"gorm.io/gorm"
)

type Album struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

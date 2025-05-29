package model

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	Name string    `gorm:"type:varchar(32)"`
}

package entity

import "residential-registration/backend/pkg/database"

type Token struct {
	ID      uint64     `gorm:"primaryKey;autoIncrement:true"`
	Value   TokenValue `gorm:"index:,unique"`
	Revoked bool
	UserID  uint64 `gorm:"index"`
	database.PostgreSQLModel
}

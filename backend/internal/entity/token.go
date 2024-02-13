package entity

import "residential-registration/backend/pkg/database"

type Token struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement:true"`
	Value      TokenValue
	Revoked    bool
	Inhabitant Inhabitant `gorm:"foreignKey:TokenID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	database.PostgreSQLModel
}

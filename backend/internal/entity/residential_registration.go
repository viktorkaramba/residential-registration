package entity

import "residential-registration/backend/pkg/database"

type FullName struct {
	FirstName  string `json:"first_name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic" binding:"required"`
}

type Inhabitant struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement:true"`
	Apartment Apartment `gorm:"foreignKey:InhabitantID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	FullName
	Password Password
	TokenID  uint64 `gorm:"index"`
	database.PostgreSQLModel
}

type Apartment struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement:true"`
	BuildingID   uint64 `gorm:"index"`
	InhabitantID uint64 `gorm:"index"`
	Number       ApartmentNumber
	Area         ApartmentArea
	database.PostgreSQLModel
}

type Building struct {
	ID         uint64      `gorm:"primaryKey;autoIncrement:true"`
	OSBBID     uint64      `gorm:"index"`
	Apartments []Apartment `gorm:"foreignKey:BuildingID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	Address    Address
	database.PostgreSQLModel
}

type OSBB struct {
	ID       uint64   `gorm:"primaryKey;autoIncrement:true"`
	Building Building `gorm:"foreignKey:OSBBID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name     Name
	EDRPOU   EDRPOU `gorm:"index"`
	database.PostgreSQLModel
}

type InputInhabitant struct {
	FullName
	Password        `json:"password" binding:"required"`
	ApartmentNumber `json:"apartment_number" binding:"required"`
	ApartmentArea   `json:"apartment_area" binding:"required"`
}

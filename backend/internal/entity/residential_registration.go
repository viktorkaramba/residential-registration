package entity

import "residential-registration/backend/pkg/database"

type FullName struct {
	FirstName  string `json:"first_name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic" binding:"required"`
}

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement:true"`
	Apartment Apartment `gorm:"foreignKey:UserID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	FullName
	Password    Password
	PhoneNumber PhoneNumber
	Token       Token  `gorm:"foreignKey:UserID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	OSBBID      uint64 `gorm:"index"`
	Role        UserRole
	database.PostgreSQLModel
}

type Apartment struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement:true"`
	BuildingID uint64 `gorm:"index"`
	UserID     uint64 `gorm:"index"`
	Number     ApartmentNumber
	Area       ApartmentArea
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
	OSBBHead User     `gorm:"foreignKey:OSBBID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	Building Building `gorm:"foreignKey:OSBBID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name     Name
	EDRPOU   EDRPOU `gorm:"index"`
	Rent     Rent
	database.PostgreSQLModel
}

type InputUser struct {
	FullName
	Password        `json:"password" binding:"required"`
	ApartmentNumber `json:"apartment_number" binding:"required"`
	ApartmentArea   `json:"apartment_area" binding:"required"`
}

type InputOSBB struct {
	FullName
	Password    `json:"password" binding:"required"`
	PhoneNumber `json:"phone_number" binding:"required"`
	Name        `json:"name" binding:"required"`
	EDRPOU      `json:"edrpou" binding:"required"`
	Address     `json:"address" binding:"required"`
	Rent        `json:"rent" binding:"required"`
}

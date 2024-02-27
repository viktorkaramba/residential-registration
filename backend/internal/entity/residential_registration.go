package entity

import (
	"residential-registration/backend/pkg/database"
	"time"
)

type FullName struct {
	FirstName  string `json:"first_name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic" binding:"required"`
}

type User struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	OSBBID uint64 `gorm:"index"`

	Apartment Apartment `gorm:"foreignKey:UserID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	FullName
	Password    Password
	PhoneNumber PhoneNumber
	Role        UserRole

	database.PostgreSQLModel
}

type Apartment struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement:true"`
	BuildingID uint64 `gorm:"index"`
	UserID     uint64 `gorm:"index"`

	Number ApartmentNumber
	Area   ApartmentArea

	database.PostgreSQLModel
}

type Building struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	OSBBID uint64 `gorm:"index"`

	Apartments []Apartment `gorm:"foreignKey:BuildingID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	Address    Address

	database.PostgreSQLModel
}

type OSBB struct {
	ID uint64 `gorm:"primaryKey;autoIncrement:true"`

	Building     Building       `gorm:"foreignKey:OSBBID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	Announcement []Announcement `gorm:"foreignKey:OSBBID;OnUpdate:CASCADE,OnDelete:CASCADE"`

	OSBBHead User `gorm:"foreignKey:OSBBID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name     Name
	EDRPOU   EDRPOU `gorm:"index"`
	Rent     Rent
	database.PostgreSQLModel
}

type Announcement struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserID uint64 `gorm:"index"`
	OSBBID uint64 `gorm:"index"`

	Title   Text
	Content Text

	CreatedAt time.Time `gorm:"index"`
	database.PostgreSQLModel
}

type Poll struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserID uint64 `gorm:"index"`
	OSBBID uint64 `gorm:"index"`

	Question    Text
	TestAnswer  []TestAnswer `gorm:"foreignKey:PollID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserAnswers []Answer     `gorm:"foreignKey:PollID;OnUpdate:CASCADE,OnDelete:CASCADE"`

	Type PollType

	CreatedAt  time.Time `gorm:"index"`
	FinishedAt time.Time `gorm:"index"`
	database.PostgreSQLModel
}

type TestAnswer struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	PollID uint64 `gorm:"index"`

	Content Text `json:"content" binding:"required"`

	database.PostgreSQLModel
}

type Answer struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement:true"`
	PollID       uint64 `gorm:"index"`
	UserID       uint64 `gorm:"index"`
	TestAnswerID uint64 `gorm:"index"`

	Content Text

	database.PostgreSQLModel
}

type Payment struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	OSBBID uint64 `gorm:"index"`

	Amount      Amount
	Appointment Appointment

	CreatedAt time.Time `gorm:"index"`
	Deadline  time.Time `gorm:"index"`

	database.PostgreSQLModel
}

type Purchase struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement:true"`
	PaymentID uint64 `gorm:"index"`
	UserID    uint64 `gorm:"index"`

	PaymentStatus PaymentStatus

	database.PostgreSQLModel
}

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
	ID           uint64         `gorm:"primaryKey;autoIncrement:true"`
	OSBBHead     User           `gorm:"foreignKey:OSBBID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	Building     Building       `gorm:"foreignKey:OSBBID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	Announcement []Announcement `gorm:"foreignKey:OSBBID;OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name         Name
	EDRPOU       EDRPOU `gorm:"index"`
	Rent         Rent
	database.PostgreSQLModel
}

type Announcement struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserID    uint64 `gorm:"index"`
	OSBBID    uint64 `gorm:"index"`
	Title     Text
	Content   Text
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

	Type       PollType
	CreatedAt  time.Time `gorm:"index"`
	FinishedAt time.Time `gorm:"index"`
	database.PostgreSQLModel
}

type TestAnswer struct {
	ID      uint64 `gorm:"primaryKey;autoIncrement:true"`
	PollID  uint64 `gorm:"index"`
	Content Text   `json:"content" binding:"required"`
	database.PostgreSQLModel
}

type Answer struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement:true"`
	PollID       uint64 `gorm:"index"`
	UserID       uint64 `gorm:"index"`
	TestAnswerID uint64 `gorm:"index"`
	Content      Text
	database.PostgreSQLModel
}

type InputUser struct {
	FirstName       string `json:"first_name" binding:"required"`
	Surname         string `json:"surname" binding:"required"`
	Patronymic      string `json:"patronymic" binding:"required"`
	Password        `json:"password" binding:"required"`
	ApartmentNumber `json:"apartment_number" binding:"required"`
	ApartmentArea   `json:"apartment_area" binding:"required"`
}

type InputOSBB struct {
	FirstName   string `json:"first_name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	Patronymic  string `json:"patronymic" binding:"required"`
	Password    `json:"password" binding:"required"`
	PhoneNumber `json:"phone_number" binding:"required"`
	Name        `json:"name" binding:"required"`
	EDRPOU      `json:"edrpou" binding:"required"`
	Address     `json:"address" binding:"required"`
	Rent        `json:"rent" binding:"required"`
}

type InputAnnouncement struct {
	Title   Text `json:"title" binding:"required"`
	Content Text `json:"content" binding:"required"`
}

type InputPoll struct {
	Question   Text      `json:"question" binding:"required"`
	FinishedAt time.Time `json:"finished_at" binding:"required"`
}

type InputPollTest struct {
	Question   Text         `json:"question" binding:"required"`
	TestAnswer []TestAnswer `json:"test_answer" binding:"required"`
	FinishedAt time.Time    `json:"finished_at" binding:"required"`
}

type InputPollAnswer struct {
	Content Text `json:"content" binding:"required"`
}

type InputPollAnswerTest struct {
	TestAnswerID uint64 `json:"test-answer-id"  binding:"required"`
}

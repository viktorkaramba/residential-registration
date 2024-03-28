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

type Apartment struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement:true" json:"id"`
	BuildingID uint64 `gorm:"index" json:"buildingID"`
	UserID     uint64 `gorm:"index" json:"userID"`

	Number ApartmentNumber `json:"number"`
	Area   ApartmentArea   `json:"area"`

	database.PostgreSQLModel
}

type User struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true" json:"-"`
	OSBBID uint64 `gorm:"index" json:"osbbid"`

	Apartment   Apartment `gorm:"foreignKey:UserID;OnUpdate:CASCADE,OnDelete:CASCADE" json:"apartment"`
	FullName    `json:"full_name"`
	Password    Password    `json:"-"`
	PhoneNumber PhoneNumber `json:"phone_number"`
	Role        UserRole    `json:"role"`
	IsApproved  *bool       `json:"is_approved"`

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
	ID uint64 `gorm:"primaryKey;autoIncrement:true" json:"id"`

	Building     Building       `gorm:"foreignKey:OSBBID;OnUpdate:CASCADE,OnDelete:CASCADE" json:"building"`
	Announcement []Announcement `gorm:"foreignKey:OSBBID;OnUpdate:CASCADE,OnDelete:CASCADE" json:"announcements"`

	OSBBHead User   `gorm:"foreignKey:OSBBID;OnUpdate:CASCADE,OnDelete:CASCADE" json:"osbb_head"`
	Name     Name   `json:"name"`
	EDRPOU   EDRPOU `gorm:"index" json:"edrpou"`
	Rent     Rent   `json:"rent"`

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
	ID     uint64 `gorm:"primaryKey;autoIncrement:true" json:"-"`
	UserID uint64 `gorm:"index" json:"-"`
	OSBBID uint64 `gorm:"index" json:"-"`

	Question    Text         `json:"question"`
	TestAnswers []TestAnswer `gorm:"foreignKey:PollID;OnUpdate:CASCADE,OnDelete:CASCADE" json:"test_answer"`
	UserAnswers []Answer     `gorm:"foreignKey:PollID;OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`

	Type PollType `json:"type"`

	CreatedAt  time.Time `gorm:"index"  json:"created_at"`
	FinishedAt time.Time `gorm:"index" json:"finished_at"`
	database.PostgreSQLModel
}

type TestAnswer struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	PollID uint64 `gorm:"index" json:"-"`

	Content Text `json:"content" binding:"required"`

	database.PostgreSQLModel
}

type Answer struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement:true"`
	PollID       uint64 `gorm:"index" json:"-"`
	UserID       uint64 `gorm:"index"`
	TestAnswerID uint64 `gorm:"index"`

	Content Text

	CreatedAt time.Time `gorm:"index"  json:"created_at"`
	UpdateAt  time.Time `gorm:"index" json:"updated_at"`

	database.PostgreSQLModel
}

type TestAnswerCount struct {
	Id    uint64 `json:"test_answer_id" db:"id"`
	Count uint64 `json:"count" db:"count"`
}

type PollResult struct {
	Answer             []Answer          `json:"answers" db:"answers"`
	CountOfTestAnswers []TestAnswerCount `json:"count_of_test_answers"`
	CountOfAllAnswers  uint64            `json:"count_of_answers"`
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

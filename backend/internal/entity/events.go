package entity

import "time"

type EventUserPayload struct {
	FirstName       string `json:"first_name" binding:"required"`
	Surname         string `json:"surname" binding:"required"`
	Patronymic      string `json:"patronymic" binding:"required"`
	Password        `json:"password" binding:"required"`
	PhoneNumber     `json:"phone_number" binding:"required"`
	ApartmentNumber `json:"apartment_number" binding:"required"`
	ApartmentArea   `json:"apartment_area" binding:"required"`
}

type EventLoginPayload struct {
	PhoneNumber `json:"phone_number" binding:"required"`
	Password    `json:"password" binding:"required"`
}

type EventOSBBPayload struct {
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

type EventAnnouncementPayload struct {
	Title   Text `json:"title" binding:"required"`
	Content Text `json:"content" binding:"required"`
}

type EventPollPayload struct {
	Question   Text      `json:"question" binding:"required"`
	FinishedAt time.Time `json:"finished_at" binding:"required"`
}

type EventPollTestPayload struct {
	Question   Text         `json:"question" binding:"required"`
	TestAnswer []TestAnswer `json:"test_answer" binding:"required"`
	FinishedAt time.Time    `json:"finished_at" binding:"required"`
}

type EventPollAnswerPayload struct {
	Content Text `json:"content" binding:"required"`
}

type EventPollAnswerTestPayload struct {
	TestAnswerID uint64 `json:"test-answer-id"  binding:"required"`
}
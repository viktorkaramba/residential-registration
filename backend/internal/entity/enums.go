package entity

type Password string

type PhoneNumber string

type Name string

type UserRole string

const (
	UserRoleInhabitant UserRole = "inhabitant"
	UserRoleOSBBHead   UserRole = "osbb_head"
)

type EDRPOU uint64

type Address string

type Rent float64

type ApartmentNumber uint64

type ApartmentArea uint64

type TokenValue string

type Text string

type PollType string

const (
	PollTypeOpenAnswer PollType = "open_answer"
	PollTypeTest       PollType = "test"
)

type Amount float64

type Appointment string

type PaymentStatus string

const (
	Paid         PaymentStatus = "paid"
	NotPaid      PaymentStatus = "not_paid"
	InProcessing PaymentStatus = "in_processing"
)

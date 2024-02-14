package entity

type Password string

type PhoneNumber string

type Name string

type UserRole string

const (
	UserRoleInhabitant UserRole = "inhabitant"
	UserRoleOSBBHEad   UserRole = "osbb_head"
)

type EDRPOU uint64

type Address string

type Rent float64

type ApartmentNumber uint64

type ApartmentArea uint64

type TokenValue string

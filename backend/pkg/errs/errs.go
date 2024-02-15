package errs

import (
	"errors"
	"log"
	"net/http"
	"strconv"
)

// Error is the type that implements the error interface.
// It contains a number of fields, each of different type.
// An Error value may leave some values unset.
type Error struct {
	// Kind is the class of error, such as permission failure,
	// or "Other" if its class is unknown or irrelevant.
	k Kind
	// Code is a human-readable, short representation of the error
	c Code
	// The underlying error that triggered this one, if any.
	err error
}

func (e *Error) isZero() bool {
	return e.k == 0 && e.c == "" && e == nil
}

// Unwrap method allows for unwrapping errors using errors.As
func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Error() string {
	return e.err.Error()
}

func M(m string) *Error {
	return Err(errors.New(m)).Kind(Other)
}

func Err(err error) *Error {
	return &Error{err: err}
}

func (e *Error) Kind(kind Kind) *Error {
	e.k = kind
	return e
}

func (e *Error) Code(code Code) *Error {
	e.c = code
	return e
}

func E(err error) *Error {
	var e *Error
	ok := errors.As(err, &e)
	if !ok {
		return &Error{
			err: err,
			k:   Other,
		}
	}

	return e
}

// TopError recursively unwraps all errors and retrieves the topmost error
func TopError(err error) error {
	currentErr := err
	for errors.Unwrap(currentErr) != nil {
		currentErr = errors.Unwrap(currentErr)
	}

	return currentErr
}

type HttpError struct {
	Code  string `json:"code"`
	Error string `json:"error,omitempty"`
}

func (e *Error) ToHttpError() HttpError {
	if e.IsServer() {
		return HttpError{
			Code:  e.c.String(),
			Error: e.err.Error(),
		}
	}

	return HttpError{
		Code: e.c.String(),
	}
}

type logger interface {
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
}

func (e *Error) Log(logger logger) {
	if e.IsServer() {
		logger.Error(e.c.String(), "err", e.err)
	} else {
		logger.Info(e.c.String(), "err", e.err)
	}
}

// Code is a human-readable, short representation of the error
type Code string

func (c Code) String() string {
	return string(c)
}

// Kind defines the kind of error this is, mostly for use by systems
// such as FUSE that must act differently depending on the error.
type Kind uint8

// Kinds of errors.
//
// The values of the error kinds are common between both
// clients and servers. Do not reorder this list or remove
// any items since that will change their values.
// New items must be added only to the end.
const (
	Other          Kind = iota // Unclassified error. This value is not printed in the error message.
	Invalid                    // Invalid operation for this type of item.
	IO                         // External I/O error such as network failure.
	Exist                      // Item already exists.
	NotExist                   // Item does not exist.
	Private                    // Information withheld.
	Internal                   // Internal error or inconsistency.
	BrokenLink                 // Link target does not exist.
	Database                   // Error from database.
	Validation                 // Input validation error.
	InvalidRequest             // Invalid Request
)

func (e *Error) String() string {
	switch e.k {
	case Other:
		return "other error"
	case Invalid:
		return "invalid operation"
	case IO:
		return "I/O error"
	case Exist:
		return "item already exists"
	case NotExist:
		return "item does not exist"
	case BrokenLink:
		return "link target does not exist"
	case Private:
		return "information withheld"
	case Internal:
		return "internal error"
	case Database:
		return "database error"
	case Validation:
		return "input validation error"
	case InvalidRequest:
		return "invalid request error"
	}
	return "unknown error kind"
}

func (e *Error) IsServer() bool {
	return e.k == Internal || e.k == Database || e.k == IO || e.k == Private || e.k == Other
}

func (e *Error) HTTPStatusCode() int {
	switch e.k {
	case Other:
		return http.StatusInternalServerError
	case Invalid:
		return http.StatusBadRequest
	case IO:
		return http.StatusServiceUnavailable
	case Exist:
		return http.StatusConflict
	case NotExist:
		return http.StatusNotFound
	case BrokenLink:
		return http.StatusGone
	case Private:
		return http.StatusForbidden
	case Internal:
		return http.StatusInternalServerError
	case Database:
		return http.StatusInternalServerError
	case Validation:
		return http.StatusBadRequest
	case InvalidRequest:
		return http.StatusBadRequest
	}
	return http.StatusUnprocessableEntity
}

// KindIs reports whether err is an *Error of the given Kind.
// If err is nil then KindIs returns false.
func KindIs(kind Kind, err error) bool {
	var e *Error
	if errors.As(err, &e) {
		if e.k != Other {
			return e.k == kind
		}
		if e != nil {
			return KindIs(kind, e)
		}
	}
	return false
}

// ParseSQLError parses sql error as string and returns error with kind.
func ParseSQLError(err error) *Error {
	var substr = []rune{'S', 'Q', 'L', 'S', 'T', 'A', 'T', 'E', '='}
	var pos = len(substr) - 1
	errRunes := []rune(err.Error())
	index := len(errRunes) - 1

	for ; index >= 0; index-- {
		if errRunes[index] == substr[pos] {
			pos--
			if pos < 0 {
				break
			}
			continue
		}
		pos = len(substr) - 1
	}

	if pos >= 0 {
		return Err(err).Kind(Database)
	}

	index += len(substr)

	var codeRunes []rune
	for ; index < len(errRunes) && errRunes[index] != ')'; index++ {
		codeRunes = append(codeRunes, errRunes[index])
	}

	code, parseErr := strconv.Atoi(string(codeRunes))
	if parseErr != nil {
		log.Println("parse error:", parseErr)
		return Err(err).Kind(Database)
	}

	switch code {
	case 23505:
		return Err(err).Kind(Exist)
	}

	return Err(err).Kind(Database)
}

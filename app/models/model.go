package models

import (
	"errors"
	"net/http"

	"github.com/golang-module/carbon/v2"
)

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id"`
}

type CommonTimestampsField struct {
	CreatedAt carbon.DateTime `gorm:"column:created_at;" json:"created_at,omitempty"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;" json:"updated_at,omitempty"`
}

type Error struct {
	Err              error
	StatusCode       int
	CustomStatusCode int
}

func (e *Error) SetError(message string, statusCode int) error {
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	e.StatusCode = statusCode
	e.Err = errors.New(message)
	return e.Err
}

func (e *Error) SetCustomStatusCode(statusCode int) *Error {
	e.CustomStatusCode = statusCode
	return e
}

func (e *Error) SetBadRequestError(message string) error {
	return e.SetError(message, http.StatusBadRequest)
}

func (e *Error) SetForbiddenError(message string) error {
	return e.SetError(message, http.StatusForbidden)
}

func (e *Error) SetNotFoundError(message string) error {
	return e.SetError(message, http.StatusNotFound)
}

func (e *Error) SetInternalError(message string) error {
	return e.SetError(message, http.StatusInternalServerError)
}

func (e *Error) SetUnauthorizedError(message string) error {
	return e.SetError(message, http.StatusUnauthorized)
}

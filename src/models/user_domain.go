package models

import (
	"github.com/google/uuid"
	"github.com/rlevidev/my-firstcrud-in-go/src/config/rest_err"
)

type UserDomain struct {
	ID       uuid.UUID
	Email    string
	Name     string
	Password string
}

type UserDomainInterface interface {
	CreateUser(UserDomain) *rest_err.RestErr
	UpdateUser(uuid.UUID, UserDomain) *rest_err.RestErr
	FindUser(uuid.UUID) (*UserDomain, *rest_err.RestErr)
	DeleteUser(uuid.UUID) *rest_err.RestErr
}

package services

import (
	"github.com/google/uuid"
	"github.com/rlevidev/my-firstcrud-in-go/src/config/rest_err"
	"github.com/rlevidev/my-firstcrud-in-go/src/models"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {

}

type UserDomainService interface {
	CreateUser(models.UserDomainInterface) *rest_err.RestErr
	UpdateUser(uuid.UUID, models.UserDomainInterface) *rest_err.RestErr
	FindUser(uuid.UUID, models.UserDomainInterface) (*models.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(uuid.UUID, models.UserDomainInterface) *rest_err.RestErr
}

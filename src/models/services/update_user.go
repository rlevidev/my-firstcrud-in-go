package services

import (
	"github.com/google/uuid"
	"github.com/rlevidev/my-firstcrud-in-go/src/config/rest_err"
	"github.com/rlevidev/my-firstcrud-in-go/src/models"
)

func (*userDomainService) UpdateUser(
	userId uuid.UUID,
	userDomain models.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}

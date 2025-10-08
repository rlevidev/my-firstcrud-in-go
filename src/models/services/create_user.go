package services

import (
	"fmt"

	"github.com/rlevidev/my-firstcrud-in-go/src/config/logger"
	"github.com/rlevidev/my-firstcrud-in-go/src/config/rest_err"
	"github.com/rlevidev/my-firstcrud-in-go/src/models"
)

func (ud *userDomainService) CreateUser(
	userDomain models.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Model de criar user iniciado.")

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	
	return nil
}

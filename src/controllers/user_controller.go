package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/my-firstcrud-in-go/src/models/services"
)

func NewUserControllerInterface(
	serviceInterface services.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{service: serviceInterface}
}

type UserControllerInterface interface {
	FindUserById(ctx *gin.Context)
	FindUserByEmail(ctx *gin.Context)

	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userControllerInterface struct {
	service services.UserDomainService
}

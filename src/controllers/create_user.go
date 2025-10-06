package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/my-firstcrud-in-go/src/config/logger"
	"github.com/rlevidev/my-firstcrud-in-go/src/config/validation"
	"github.com/rlevidev/my-firstcrud-in-go/src/controllers/models/request"
)

func CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro ao tentar validar a solicitação do usuário", err)
		errRest := validation.ValidateUserError(err)
		
		ctx.JSON(errRest.Status, errRest)
	}
}

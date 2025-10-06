package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rlevidev/my-firstcrud-in-go/src/config/validation"
	"github.com/rlevidev/my-firstcrud-in-go/src/controllers/models/request"
)

func CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Erro ao tentar serializar o objeto, erro: %s", err.Error())
		errRest := validation.ValidateUserError(err)

		ctx.JSON(errRest.Status, errRest)
	}
}

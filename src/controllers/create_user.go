package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlevidev/my-firstcrud-in-go/src/config/logger"
	"github.com/rlevidev/my-firstcrud-in-go/src/config/validation"
	"github.com/rlevidev/my-firstcrud-in-go/src/controllers/model/request"
	"github.com/rlevidev/my-firstcrud-in-go/src/models"
	"github.com/rlevidev/my-firstcrud-in-go/src/view"
)

var (
	UserDomainInterface models.UserDomainInterface
)

func (uci *userControllerInterface) CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro ao tentar validar a solicitação do usuário", err)
		errRest := validation.ValidateUserError(err)

		ctx.JSON(errRest.Status, errRest)
	}

	domain := models.NewUserDomain(
		userRequest.Email,
		userRequest.Name,
		userRequest.Password,
	)
	if err := uci.service.CreateUser(domain); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	logger.Info("Usuário criado com sucesso")
	ctx.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domain,
	))
}

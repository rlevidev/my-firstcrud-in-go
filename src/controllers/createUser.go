package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rlevidev/my-firstcrud-in-go/src/config/rest_err"
	"github.com/rlevidev/my-firstcrud-in-go/src/controllers/models/request"
)

func CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Há alguns campos incorretos em sua requisição. Error: %s", err.Error()), nil)
		ctx.JSON(restErr.Status, restErr)
	}
}

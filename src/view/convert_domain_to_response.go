package view

import (
	"github.com/rlevidev/my-firstcrud-in-go/src/controllers/model/response"
	"github.com/rlevidev/my-firstcrud-in-go/src/models"
)

func ConvertDomainToResponse(userDomain models.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetId().String(),
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
	}
}

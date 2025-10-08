package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/my-firstcrud-in-go/src/controllers"
)

func InitRoutes(
	rout *gin.RouterGroup,
	userController controllers.UserControllerInterface) {

	v1 := rout.Group("/api/v1")

	v1.GET("/getUserById/:userId", userController.FindUserById)
	v1.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	v1.POST("/createUser", userController.CreateUser)
	v1.PUT("/updateUser/:userId", userController.UpdateUser)
	v1.DELETE("/deleteUser/:userId", userController.DeleteUser)
}

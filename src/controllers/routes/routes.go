package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/my-firstcrud-in-go/src/controllers"
)

func InitRoutes(rout *gin.RouterGroup) {
	v1 := rout.Group("/api/v1")

	v1.GET("/getUserById/:userId", controllers.FindUserById)
	v1.GET("/getUserByEmail/:userEmail", controllers.FindUserByEmail)
	v1.POST("/createUser", controllers.CreateUser)
	v1.PUT("/updateUser/:userId", controllers.UpdateUser)
	v1.DELETE("/deleteUser/:userId", controllers.DeleteUser)
}

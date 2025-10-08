package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/my-firstcrud-in-go/src/controllers"
	"github.com/rlevidev/my-firstcrud-in-go/src/controllers/routes"
	"github.com/rlevidev/my-firstcrud-in-go/src/models/services"
)

func main() {
	router := gin.Default()

	service := services.NewUserDomainService()
	userController := controllers.NewUserControllerInterface(service)

	routes.InitRoutes(&router.RouterGroup, userController)
	router.Run(":8080")
}

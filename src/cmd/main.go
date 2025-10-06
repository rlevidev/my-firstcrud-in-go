package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/my-firstcrud-in-go/src/controllers/routes"
)

func main() {
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)
	router.Run(":8080")
}

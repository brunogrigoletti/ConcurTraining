package main

import (
	"github.com/gin-gonic/gin"
	"apiClientes/endpoints"
)

func main() {
	router := gin.Default()
	router.GET("/clients", endpoints.GetAllUsers)
	router.GET("/clients/:id", endpoints.GetUserById)
	router.POST("/newclient", endpoints.PostClient)
	router.Run("localhost:8080")
}
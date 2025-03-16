package main

import (
	"apiClientes/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/clients", endpoints.GetAllUsers)
	router.GET("/clients/:id", endpoints.GetUserById)
	router.POST("/newclient", endpoints.PostClient)
	router.PUT("/updateclient/:id", endpoints.UpdateClient)
	router.DELETE("/deleteclient/:id", endpoints.DeleteClient)
	router.Run("localhost:8080")
}

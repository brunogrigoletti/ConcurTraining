package endpoints

import (
	"apiClientes/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, structs.Clients)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	for _, client := range structs.Clients {
		if client.ID == id {
			c.IndentedJSON(http.StatusOK, client)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Client not found!"})
}

func PostClient(c *gin.Context) {
	var newClient structs.Client
	if err := c.BindJSON(&newClient); err != nil {
		return
	}
	structs.Clients = append(structs.Clients, newClient)
	c.IndentedJSON(http.StatusCreated, newClient)
}

func DeleteClient(c *gin.Context) {
	id := c.Param("id")
	for i, client := range structs.Clients {
		if client.ID == id {
			structs.Clients = append(structs.Clients[:i], structs.Clients[i+1:]...) //Go não tem um método de remoção de elementos de um slice, então é necessário adicionar uma nova slice removendo o elemento desejado
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Client deleted!"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Client not found!"})
}

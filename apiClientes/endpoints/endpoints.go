package endpoints

import (
	"database/sql"
	"log"
	"apiClientes/structs"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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

func PostClient(db *sql.DB) {
	var newClient structs.Client
	if err := c.BindJSON(&newClient); err != nil {
		return
	}

	stmt, err := db.Prepare("INSERT INTO clients (id, name, birthdate, email) VALUES ('9','Rafael Godoy','10-10-1991','rafael.godoy@sap.com')")
	if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
	}
	defer stmt.Close()
	_, err = stmt.Exec(produto.Nome)
	if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
	}
	//structs.Clients = append(structs.Clients, newClient)
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

func UpdateClient(c *gin.Context) { //tem que mandar todas as informações do cliente, se não o que não mandar fica vazio
	id := c.Param("id")

	var updatedClient structs.Client
	if err := c.BindJSON(&updatedClient); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request!"})
		return
	}

	for i, client := range structs.Clients {
		if client.ID == id {
			structs.Clients[i] = updatedClient
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Client updated!"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Client not found!"})
}
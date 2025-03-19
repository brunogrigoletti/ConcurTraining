package endpoints

import (
	"database/sql"
	"apiClientes/structs"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetAllUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM clients")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var clients []structs.Client
		for rows.Next() {
			var client structs.Client
			err := rows.Scan(&client.ID, &client.Name, &client.Birthdate, &client.Email)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			clients = append(clients, client)
		}
		c.JSON(http.StatusOK, clients)
	}
}

func GetUserById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		row := db.QueryRow("SELECT * FROM clients WHERE id = $1", id)

		var client structs.Client
		err := row.Scan(&client.ID, &client.Name, &client.Birthdate, &client.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, client)
	}
}

func PostClient(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var newClient structs.Client
		if err := c.BindJSON(&newClient); err != nil {
			return
		}

		stmt, err := db.Prepare("INSERT INTO clients (id, name, birthdate, email) VALUES ($1, $2, $3, $4)")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(newClient.ID, newClient.Name, newClient.Birthdate, newClient.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusCreated, newClient)
	}
}

/* func DeleteClient(c *gin.Context) {
	id := c.Param("id")
	for i, client := range structs.Clients {
		if client.ID == id {
			structs.Clients = append(structs.Clients[:i], structs.Clients[i+1:]...) //Go não tem um método de remoção de elementos de um slice, então é necessário adicionar uma nova slice removendo o elemento desejado
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Client deleted!"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Client not found!"})
} */

/* func UpdateClient(c *gin.Context) { //tem que mandar todas as informações do cliente, se não o que não mandar fica vazio
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
} */
package main

import (
	"database/sql"
	"log"
	"apiClientes/endpoints"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=postgres dbname=clients_db sslmode=disable")
 	if err != nil {
  		log.Fatal(err)
 	}
 	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS clients
		(
			id SERIAL PRIMARY KEY,
			nome VARCHAR(255) NOT NULL,
			birthdate VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL
		)`)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/clients", endpoints.GetAllUsers)
	router.GET("/clients/:id", endpoints.GetUserById)
	router.POST("/newclient", endpoints.PostClient)
	router.PUT("/updateclient/:id", endpoints.UpdateClient)
	router.DELETE("/deleteclient/:id", endpoints.DeleteClient)
	router.Run("localhost:8080")
}
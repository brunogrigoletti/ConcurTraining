package structs

type Client struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	Email     string `json:"email"`
}

var Clients = []Client{
	{ID: "1", Name: "Bruno Laitano", Birthdate: "04-14-1997", Email: "bruno.laitano@sap.com"},
	{ID: "2", Name: "Thiago Fanfa", Birthdate: "02-07-1995", Email: "thiago.fanfa@sap.com"},
	{ID: "3", Name: "Fabricio Schuh", Birthdate: "02-20-1999", Email: "fabricio.schuh@sap.com"},
}

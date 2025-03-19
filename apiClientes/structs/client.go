package structs
type Client struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	Email     string `json:"email"`
}
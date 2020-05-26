package domain

// struct User for parsing login credentials
type User struct {
	ID           uint64 `json: "id"`
	FirstName    string `json: "firstname"`
	LastName     string `json: "lastname"`
	Email        string `json: "email"`
	Password     string `json: "password"`
	HashPassword []byte `json: "hashpassword", omitempty`
}

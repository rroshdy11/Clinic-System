package entities

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Type     string `json:"type"`
}

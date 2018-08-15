package models

// User handles user data
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JwtToken handles issued token
type JwtToken struct {
	Token string `json:"token"`
}

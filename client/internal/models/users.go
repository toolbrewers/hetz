package models

type CreateUser struct {
	Username     string
	Email        string
	Password     string
	HetznerToken string
}

type GetUser struct {
	ID       uint64
	Email    string
	Username string
}

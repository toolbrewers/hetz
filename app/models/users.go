package models

type CreateUser struct {
	Username     string
	Email        string
	Password     string
	HetznerToken string
}

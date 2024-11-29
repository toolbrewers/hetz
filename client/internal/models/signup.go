package models

import "github.com/go-playground/validator/v10"

var SignupValidations = map[string]func(
	validate *validator.Validate,
	field string,
) error{
	"username": func(validate *validator.Validate, field string) error {
		return validate.Var(field, "alphanum,gt=2,lte=32")
	},
	"email": func(validate *validator.Validate, field string) error {
		return validate.Var(field, "email")
	},
	"password": func(validate *validator.Validate, field string) error {
		return validate.Var(field, "gt=8,lte=72")
	},
	"hetzner_token": func(validate *validator.Validate, field string) error {
		return validate.Var(field, "eq=64")
	},
}

var SignupHelpers = map[string]string{
	"username":      "Username's length must be greater than two and contain only alphanumeric characters.",
	"email":         "Email must be a valid address.",
	"password":      "Password's length must be greater than eight.",
	"hetzner_token": "Hetzner's token must be a valid Hetzner Cloud API token.",
}

type Signup struct {
	Username     string `form:"username"`
	Email        string `form:"email"`
	Password     string `form:"password"`
	HetznerToken string `form:"hetzner_token"`
}

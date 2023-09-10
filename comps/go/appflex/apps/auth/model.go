package auth

import (
	"time"
)

type Status string

func (u Status) String() string {
	return string(u)
}

type Query struct {
	Id    string
	Email string
}

// variables to store public and private routers
var publicRouters = map[string]string{}
var adminRouters = map[string]string{}

const (
	StatusActive      Status = "active"
	StatusUnconfirmed Status = "unconfirmed"
)

// StatusAll variable with all the status converted into string
var StatusAll = [4]string{
	StatusActive.String(),
	StatusUnconfirmed.String(),
}

// NewModel factory to Create mew struct user model
func NewModel(name, email string) Model {
	return Model{Name: name, Email: email}
}

// Model structure to represent a user on the service, this model has to represent everything related to the user
type Model struct {
	Id               string    `json:"id"`
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
	Status           Status    `json:"status,omitempty"`
	Name             string    `json:"name,omitempty"`
	Password         string    `json:"password,omitempty"`
	Email            string    `json:"email,omitempty"`
	Age              uint      `json:"age,omitempty"`
	ConfirmationCode string    `json:"confirmationCode,omitempty"`
}

func (Model) TableName() string {
	return "users"
}

// loginInfoModel estrutura usada para ser usada no retorno do login
type loginInfoModel struct {
	Status Status
	UserId string `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

// loginModel struct that represent login
type loginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

package services

import "github.com/felipdc/arqbusca-api/models"

type UserService interface {
	CreateUser(*models.User) (*models.User, error)
	GetUser(*string) (*models.User, error)
	GetAll() (*models.User, error)
}

package repository

import (
	"context"
	"rest-websockets/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error          //insertar
	GetUserById(ctx context.Context, id string) (*models.User, error) //devolver un user y un error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	Close() error
}

var implementacion UserRepository

//set repository
func SetRepository(repository UserRepository) {
	implementacion = repository
}

//insertar nuevos usuarios
func InsertUser(ctx context.Context, user *models.User) error {
	return implementacion.InsertUser(ctx, user)
}

//traer usuario conforme su id
func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementacion.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementacion.GetUserByEmail(ctx, email)
}

func Close() error {
	return implementacion.Close()
}

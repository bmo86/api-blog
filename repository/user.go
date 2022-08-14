package repository

import (
	"context"
	"rest-websockets/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error         //insertar
	GetUserById(ctx context.Context, id int64) (*models.User, error) //devolver un user y un error
}

var implementacion UserRepository

func SetRepository(repo UserRepository) {
	implementacion = repo
}

//insertar nuevos usuarios
func InsertUser(ctx context.Context, user *models.User) error {
	return implementacion.InsertUser(ctx, user)
}

//traer usuario conforme su id
func GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return implementacion.GetUserById(ctx, id)
}

package repository

import (
	"context"
	"rest-websockets/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error          //insertar
	GetUserById(ctx context.Context, id string) (*models.User, error) //devolver un user y un error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertPost(ctx context.Context, post *models.Post) error
	Close() error
}

var implementacion Repository

//set repository
func SetRepository(repository Repository) {
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

//obtener el usuario con el email
func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementacion.GetUserByEmail(ctx, email)
}

//insertar un nuevo post
func InsertPost(ctx context.Context, post *models.Post) error {
	return implementacion.InsertPost(ctx, post)
}

func Close() error {
	return implementacion.Close()
}

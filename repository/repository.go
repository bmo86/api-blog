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
	GetPostById(ctx context.Context, id string) (*models.Post, error)
	UpdatePost(ctx context.Context, p *models.Post) error
	DeletePost(ctx context.Context, id string, userId string) error
	ListPost(ctx context.Context, page uint64) ([]*models.Post, error)
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

//traer post con el id
func GetPostById(ctx context.Context, id string) (*models.Post, error) {
	return implementacion.GetPostById(ctx, id)
}

func UpdatePost(ctx context.Context, p *models.Post) error {
	return implementacion.UpdatePost(ctx, p)
}

func DeletePost(ctx context.Context, id string, userId string) error {
	return implementacion.DeletePost(ctx, id, userId)
}

func ListPost(ctx context.Context, page uint64) ([]*models.Post, error) {
	return implementacion.ListPost(ctx, page)
}

func Close() error {
	return implementacion.Close()
}

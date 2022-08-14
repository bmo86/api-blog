package database

import (
	"context"
	"database/sql"
	"log"
	"rest-websockets/models"

	_ "github.com/lib/pq"
)

type PostgraesRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgraesRepository, error) {
	db, err := sql.Open("postgres", url) //abrir repo y indicar db
	if err != nil {
		return nil, err
	}
	return &PostgraesRepository{db}, nil
}

func (repo *PostgraesRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password) //ejecutar una oracio de sql
	return err
}

func (repo *PostgraesRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *PostgraesRepository) CLose() error {
	return repo.db.Close()
}

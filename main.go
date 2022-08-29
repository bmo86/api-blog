package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"rest-websockets/handlres"
	"rest-websockets/middleware"
	"rest-websockets/server"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env") // cargar archivo

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)

}

func BindRoutes(s server.Server, r *mux.Router) {

	api := r.PathPrefix("/api/v1").Subrouter()

	api.Use(middleware.CheckAuthMiddleware(s))
	r.HandleFunc("/", handlres.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/singup", handlres.SingUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlres.LoginHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/me", handlres.MeHandler(s)).Methods(http.MethodGet)
	api.HandleFunc("/posts", handlres.InsertPostHandlres(s)).Methods(http.MethodPost)
	api.HandleFunc("/posts/{postId}", handlres.GetPostByIdHandler(s)).Methods(http.MethodGet)
	api.HandleFunc("/posts/{postId}", handlres.UpdatePostHandler(s)).Methods(http.MethodPut)
	api.HandleFunc("/posts/{postId}", handlres.DeletePostHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/posts", handlres.ListPostHandler(s)).Methods(http.MethodGet)

	
	r.HandleFunc("/ws", s.Hub().HandleWebSocket)
}

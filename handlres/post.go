package handlres

import (
	"encoding/json"
	"net/http"
	"rest-websockets/means"
	"rest-websockets/models"
	"rest-websockets/repository"
	"rest-websockets/server"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

type UpsertPostResquest struct {
	PostContent string `json:"post_content"`
}

type PostResponse struct {
	Id          string `json:"id"`
	PostContent string `json:"post_content"`
}

type MsgResponse struct {
	Msg string `json:"msg"`
}

func InsertPostHandlres(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token, err := means.Token(s, w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*models.AppClaimas); ok && token.Valid {
			var postRequest = UpsertPostResquest{}
			if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			post := models.Post{
				Id:          id.String(),
				PostContent: postRequest.PostContent,
				UserId:      claims.UserId,
			}

			err = repository.InsertPost(r.Context(), &post)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostResponse{
				Id:          post.Id,
				PostContent: post.PostContent,
			})

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func GetPostByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// -> /post/{id}
		params := mux.Vars(r) //estraer id

		posts, err := repository.GetPostById(r.Context(), params["postId"])

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(posts.Id) <= 0 {
			http.Error(w, "post not found "+err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	}

}

func UpdatePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		token, err := means.Token(s, w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*models.AppClaimas); ok && token.Valid {
			var postRequest = UpsertPostResquest{}
			if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			updatePost := models.Post{
				Id:          params["postId"],
				PostContent: postRequest.PostContent,
				UserId:      claims.UserId,
			}

			err = repository.UpdatePost(r.Context(), &updatePost)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(MsgResponse{
				Msg: "Post Update",
			})

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}

func DeletePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		token, err := means.Token(s, w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*models.AppClaimas); ok && token.Valid {

			err = repository.DeletePost(r.Context(), params["postId"], claims.UserId)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(MsgResponse{
				Msg: "Post Delete!",
			})

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}

func ListPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		/*
			/posts?page=example
		*/
		pageStr := r.URL.Query().Get("page")
		var page = uint64(0)
		if pageStr != "" {
			page, err = strconv.ParseUint(pageStr, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		posts, err := repository.ListPost(r.Context(), page)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	}
}

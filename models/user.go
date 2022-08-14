package models

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"string"`
	Password string `json:"pass"`
}

package models

type User struct {
	Id   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

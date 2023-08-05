package models

type User struct {
	Id string `json:"id"`
	Age  int64  `json:"age"`
	Name string `json:"name"`
	Description string `json:"description"`
}

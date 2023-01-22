package models

type User struct {
	Name     string `redis:"name"`
	LastName string `redis:"lastName"`
	Age      int    `redis:"age"`
}

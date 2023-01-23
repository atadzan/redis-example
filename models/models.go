package models

type User struct {
	Name     string `redis:"name"`
	LastName string `redis:"lastName"`
	Age      int    `redis:"age"`
	Sub      Video  `redis:"sub"`
}

type Video struct {
	Id   int    `redis:"id"`
	Name string `redis:"name"`
}

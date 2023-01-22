package main

type user struct {
	Name     string
	Id       int
	Password string `json:"password"`
}

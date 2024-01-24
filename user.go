package main

type User string

func NewUser(name string) User {
	return User(name)
}

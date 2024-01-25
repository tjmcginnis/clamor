package main

import "math/rand"

var colors = []string{
	"red",
	"yellow",
	"green",
	"blue",
	"indigo",
	"purple",
	"pink",
}

type User struct {
	Name  string
	Color string
}

func NewUser(name string) *User {
	color := colors[rand.Intn(len(colors))]
	return &User{
		Name:  name,
		Color: color,
	}
}

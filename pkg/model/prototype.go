package model

import "github.com/rafaelsanzio/go-core/pkg/user"

func PrototypeUser() user.User {
	return user.User{
		Name: "John Doe",
		Age:  38,
	}
}

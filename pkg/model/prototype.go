package model

import "github.com/rafaelsanzio/go-core/pkg/user"

func PrototypeUser() user.User {
	return user.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Username:  "johndoe",
		Email:     "john@mail.com",
	}
}

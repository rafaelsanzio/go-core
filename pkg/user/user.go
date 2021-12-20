package user

import (
	"context"

	"github.com/rafaelsanzio/go-core/pkg/errs"
)

type UserRepo interface {
	Insert(ctx context.Context, u User) errs.AppError
	GetUser(ctx context.Context, id string) (*User, errs.AppError)
}

type User struct {
	ID   string `bson:"_id"`
	Name string
	Age  int
}

func (u User) GetID() string {
	return u.ID
}

func (u *User) SetID(id string) {
	u.ID = id
}

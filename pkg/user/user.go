package user

import (
	"context"

	"github.com/rafaelsanzio/go-core/pkg/errs"
)

type UserRepo interface {
	Insert(ctx context.Context, u User) errs.AppError
	Get(ctx context.Context, id string) (*User, errs.AppError)
	List(ctx context.Context) ([]User, errs.AppError)
	Update(ctx context.Context, u User) (*User, errs.AppError)
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

package user

import (
	"context"
	"time"

	"github.com/rafaelsanzio/go-core/pkg/errs"
)

type UserRepo interface {
	Insert(ctx context.Context, u User) errs.AppError
	Get(ctx context.Context, id string) (*User, errs.AppError)
	List(ctx context.Context) ([]User, errs.AppError)
	Update(ctx context.Context, u User) (*User, errs.AppError)
	Delete(ctx context.Context, id string) errs.AppError
}

type User struct {
	ID        string `bson:"_id"`
	FirstName string
	LastName  string
	Username  string
	Email     string
	CreatedAt time.Time
}

func (u User) GetID() string {
	return u.ID
}

func (u *User) SetID(id string) {
	u.ID = id
}

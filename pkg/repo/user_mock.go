package repo

import (
	"context"
	"fmt"

	"github.com/rafaelsanzio/go-core/pkg/errs"
	"github.com/rafaelsanzio/go-core/pkg/user"
)

type MockUserRepo struct {
	user.UserRepo
	InsertFunc  func(ctx context.Context, u user.User) errs.AppError
	GetUserFunc func(ctx context.Context, id string) (*user.User, errs.AppError)
}

func (m MockUserRepo) Insert(ctx context.Context, u user.User) errs.AppError {
	if m.InsertFunc != nil {
		return m.InsertFunc(ctx, u)
	}
	fmt.Println("InsertFunc is nil", u)
	return m.UserRepo.Insert(ctx, u)
}

func (m MockUserRepo) GetUser(ctx context.Context, id string) (*user.User, errs.AppError) {
	if m.GetUserFunc != nil {
		return m.GetUserFunc(ctx, id)
	}
	return m.UserRepo.GetUser(ctx, id)
}

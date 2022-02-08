package repo

import (
	"context"

	"github.com/rafaelsanzio/go-core/pkg/errs"
	"github.com/rafaelsanzio/go-core/pkg/user"
)

type MockUserRepo struct {
	user.UserRepo
	InsertFunc func(ctx context.Context, u user.User) errs.AppError
	GetFunc    func(ctx context.Context, id string) (*user.User, errs.AppError)
	ListFunc   func(ctx context.Context) ([]user.User, errs.AppError)
	UpdateFunc func(ctx context.Context, u user.User) (*user.User, errs.AppError)
	DeleteFunc func(ctx context.Context, id string) errs.AppError
}

func (m MockUserRepo) Insert(ctx context.Context, u user.User) errs.AppError {
	if m.InsertFunc != nil {
		return m.InsertFunc(ctx, u)
	}
	return m.UserRepo.Insert(ctx, u)
}

func (m MockUserRepo) Get(ctx context.Context, id string) (*user.User, errs.AppError) {
	if m.GetFunc != nil {
		return m.GetFunc(ctx, id)
	}
	return m.UserRepo.Get(ctx, id)
}

func (m MockUserRepo) List(ctx context.Context) ([]user.User, errs.AppError) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx)
	}
	return m.UserRepo.List(ctx)
}

func (m MockUserRepo) Update(ctx context.Context, u user.User) (*user.User, errs.AppError) {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, u)
	}
	return m.UserRepo.Update(ctx, u)
}

func (m MockUserRepo) Delete(ctx context.Context, id string) errs.AppError {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, id)
	}
	return m.UserRepo.Delete(ctx, id)
}

package repo

import (
	"context"
	"testing"

	"github.com/rafaelsanzio/go-core/pkg/errs"
	"github.com/rafaelsanzio/go-core/pkg/model"
	"github.com/rafaelsanzio/go-core/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestUserRepoInsert(t *testing.T) {
	ctx := context.Background()

	SetUserRepo(MockUserRepo{
		InsertFunc: func(ctx context.Context, u user.User) errs.AppError {
			return nil
		},
	})
	defer SetUserRepo(nil)

	newUser := model.PrototypeUser()

	err := GetUserRepo().Insert(ctx, newUser)
	assert.NoError(t, err)
}

func TestUserRepoGet(t *testing.T) {
	ctx := context.Background()

	SetUserRepo(MockUserRepo{
		GetFunc: func(ctx context.Context, id string) (*user.User, errs.AppError) {
			user := model.PrototypeUser()
			return &user, nil
		},
	})
	defer SetUserRepo(nil)

	newUser := model.PrototypeUser()

	user, err := GetUserRepo().Get(ctx, "new-user-id")
	assert.NoError(t, err)

	assert.Equal(t, newUser.Name, user.Name)
	assert.Equal(t, newUser.Age, user.Age)
}

func TestUserRepoList(t *testing.T) {
	ctx := context.Background()

	SetUserRepo(MockUserRepo{
		ListFunc: func(ctx context.Context) ([]user.User, errs.AppError) {
			userMock := model.PrototypeUser()
			userMock2 := model.PrototypeUser()

			return []user.User{userMock, userMock2}, nil
		},
	})
	defer SetUserRepo(nil)

	users, err := GetUserRepo().List(ctx)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(users))
}

func TestUserRepoUpdate(t *testing.T) {
	ctx := context.Background()

	SetUserRepo(MockUserRepo{
		UpdateFunc: func(ctx context.Context, u user.User) (*user.User, errs.AppError) {
			return &u, nil
		},
	})
	defer SetUserRepo(nil)

	newUser := model.PrototypeUser()

	userUpdated, err := GetUserRepo().Update(ctx, newUser)
	assert.NoError(t, err)

	assert.Equal(t, newUser.Name, userUpdated.Name)
	assert.Equal(t, newUser.Age, userUpdated.Age)
}

func TestUserRepoDelete(t *testing.T) {
	ctx := context.Background()

	SetUserRepo(MockUserRepo{
		DeleteFunc: func(ctx context.Context, id string) errs.AppError {
			return nil
		},
	})
	defer SetUserRepo(nil)

	newUser := model.PrototypeUser()

	err := GetUserRepo().Delete(ctx, newUser.GetID())
	assert.NoError(t, err)
}

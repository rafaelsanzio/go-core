package repo

import (
	"context"
	"fmt"

	"github.com/rafaelsanzio/go-core/pkg/applog"
	"github.com/rafaelsanzio/go-core/pkg/errs"
	"github.com/rafaelsanzio/go-core/pkg/store"
	"github.com/rafaelsanzio/go-core/pkg/store/query"
	"github.com/rafaelsanzio/go-core/pkg/user"
)

const (
	UserCollection = "user"
)

type userRepo struct {
	store store.Store
}

var userRepoSingleton user.UserRepo

func GetUserRepo() user.UserRepo {
	if userRepoSingleton == nil {
		return getUserRepo()
	}
	return userRepoSingleton
}

func getUserRepo() *userRepo {
	s := store.GetStore()
	return &userRepo{s}
}

func SetUserRepo(repo user.UserRepo) {
	userRepoSingleton = repo
}

func (repo userRepo) Insert(ctx context.Context, u user.User) errs.AppError {
	fmt.Println("Inserting user: ", u)
	_, err := repo.store.InsertOne(ctx, UserCollection, &u)
	return err
}

func (repo userRepo) GetUser(ctx context.Context, id string) (user.User, errs.AppError) {
	filter := query.Filter{
		"id": id,
	}

	fmt.Println("Getting user: ", id)

	opts := query.FindOneOptions{}

	mUser := user.User{}
	err := repo.store.FindOne(ctx, UserCollection, filter, &mUser, opts)
	if err != nil {
		return mUser, errs.ErrMongoFindOne.Throwf(applog.Log, "for collection: %s, and ID: %s, err: [%v]", UserCollection, id, err)
	}

	return mUser, nil
}

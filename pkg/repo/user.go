package repo

import (
	"context"

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
	_, err := repo.store.InsertOne(ctx, UserCollection, &u)
	return err
}

func (repo userRepo) Get(ctx context.Context, id string) (*user.User, errs.AppError) {
	filter := query.Filter{
		"id": id,
	}

	opts := query.FindOneOptions{}

	mUser := user.User{}
	err := repo.store.FindOne(ctx, UserCollection, filter, &mUser, opts)
	if err != nil {
		return &mUser, errs.ErrMongoFindOne.Throwf(applog.Log, "for collection: %s, and ID: %s, err: [%v]", UserCollection, id, err)
	}

	return &mUser, nil
}

func (repo userRepo) List(ctx context.Context) ([]user.User, errs.AppError) {
	filter := query.Filter{}

	opts := query.FindOptions{}
	mUser := []user.User{}
	users, err := repo.store.Find(ctx, UserCollection, filter, opts)
	if err != nil {
		return mUser, errs.ErrMongoFind.Throwf(applog.Log, "for collection: %s, err: [%v]", UserCollection, err)
	}

	defer func() {
		_ = users.Close(ctx)
	}()

	for {
		if users.Err() != nil {
			return mUser, err
		}

		if ok := users.Next(ctx); !ok {
			break
		}

		var u user.User
		if err_ := users.Decode(&u); err_ != nil {
			return mUser, err
		}

		mUser = append(mUser, u)
	}

	return mUser, nil
}

func (repo userRepo) Update(ctx context.Context, u user.User) (*user.User, errs.AppError) {
	res := user.User{}
	filter := query.Filter{
		"id": u.GetID(),
	}

	err := repo.store.FindOne(ctx, UserCollection, filter, &res)
	if err != nil {
		return nil, err
	}

	if res.ID == "" {
		return nil, errs.ErrMongoFindOne.Throwf(applog.Log, "for collection: %s, and ID: %s, err: [%v]", UserCollection, u.GetID(), err)
	}

	u.ID = res.ID
	err = repo.store.UpdateOne(ctx, UserCollection, &u)
	if err != nil {
		return nil, errs.ErrMongoUpdateOne.Throwf(applog.Log, "for collection: %s, and ID: %s, err: [%v]", UserCollection, u.GetID(), err)
	}

	return &u, nil
}

func (repo userRepo) Delete(ctx context.Context, id string) errs.AppError {
	err := repo.store.DeleteOne(ctx, UserCollection, id)
	return err
}

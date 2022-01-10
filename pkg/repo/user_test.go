package repo

import (
	"context"
	"testing"

	"github.com/rafaelsanzio/go-core/pkg/model"
	"github.com/rafaelsanzio/go-core/pkg/store"

	"github.com/stretchr/testify/assert"
)

func TestUserRepoInsert(t *testing.T) {
	ctx := context.Background()
	newUser := model.PrototypeUser()

	err := GetUserRepo().Insert(ctx, newUser)
	assert.NoError(t, err)
	_ = store.GetStore().DeleteOne(ctx, UserCollection, newUser.ID)
}

func TestUserRepoGet(t *testing.T) {
	ctx := context.Background()
	newUser := model.PrototypeUser()

	err := GetUserRepo().Insert(ctx, newUser)
	assert.NoError(t, err)

	user, err := GetUserRepo().Get(ctx, newUser.ID)
	assert.NoError(t, err)

	assert.Equal(t, newUser.ID, user.ID)
	assert.Equal(t, newUser.Name, user.Name)
	assert.Equal(t, newUser.Age, user.Age)

	_ = store.GetStore().DeleteOne(ctx, UserCollection, newUser.ID)
}

func TestUserRepoList(t *testing.T) {
	ctx := context.Background()
	newUser := model.PrototypeUser()

	err := GetUserRepo().Insert(ctx, newUser)
	assert.NoError(t, err)
	err = GetUserRepo().Insert(ctx, newUser)
	assert.NoError(t, err)

	users, err := GetUserRepo().List(ctx)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(users))

	_ = store.GetStore().DeleteOne(ctx, UserCollection, users[0].ID)
	_ = store.GetStore().DeleteOne(ctx, UserCollection, users[1].ID)
}

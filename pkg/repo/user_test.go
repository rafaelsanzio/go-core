package repo

import (
	"context"
	"fmt"
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

func TestUserRepoGetUser(t *testing.T) {
	ctx := context.Background()
	newUser := model.PrototypeUser()

	err := GetUserRepo().Insert(ctx, newUser)
	assert.NoError(t, err)

	user, err := GetUserRepo().GetUser(ctx, newUser.ID)
	fmt.Println("user: ", user)
	assert.NoError(t, err)

	assert.Equal(t, newUser.ID, user.ID)
	assert.Equal(t, newUser.Name, user.Name)
	assert.Equal(t, newUser.Age, user.Age)

	_ = store.GetStore().DeleteOne(ctx, UserCollection, newUser.ID)
}

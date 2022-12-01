package loader

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/repository"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/Sei-Yukinari/gqlgen-todos/src/util/apperror"
	"github.com/graph-gophers/dataloader"
)

type UserLoader struct {
	userRepository repository.User
}

func (u *UserLoader) BatchGetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	userIDs := make([]int, len(keys))
	for i, key := range keys {
		userIDs[i], _ = strconv.Atoi(key.String())
	}
	logger.Debugf("BatchGetUsers(id = %v)\n", userIDs)
	users, err := u.userRepository.FindByIDs(ctx, userIDs)
	if err != nil {
		err := fmt.Errorf("fail get users, %w", err)
		logger.Fatalf("%v\n", err)
		return nil
	}

	output := make([]*dataloader.Result, len(keys))
	for index, v := range users {
		output[index] = &dataloader.Result{Data: v, Error: nil}
	}
	return output
}

func LoadUser(ctx context.Context, userID string) (*model.User, apperror.AppError) {
	l := logger.FromContext(ctx)
	l.Debugf("LoadUser(id = %s)\n", userID)
	loaders, err := FromContext(ctx)
	if err != nil {
		return nil, apperror.Wrap(err)
	}
	thunk := loaders.UserLoader.Load(ctx, dataloader.StringKey(userID))
	result, e := thunk()
	if err != nil {
		return nil, apperror.Wrap(e)
	}
	user := result.(*model.User)
	logger.Debugf("return LoadUser(id = %d, name = %s)\n", user.ID, user.Name)
	return user, nil
}

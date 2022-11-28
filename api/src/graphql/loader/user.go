package loader

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/repository"
	"github.com/graph-gophers/dataloader"
)

type UserLoader struct {
	userRepository repository.User
}

func (u *UserLoader) BatchGetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// dataloader.Keysの型を[]stringに変換する
	userIDs := make([]int, len(keys))
	for i, key := range keys {
		userIDs[i], _ = strconv.Atoi(key.String())
	}
	// 実処理
	log.Printf("BatchGetUsers(id = %v)\n", userIDs)
	users, err := u.userRepository.FindByIDs(ctx, userIDs)
	if err != nil {
		err := fmt.Errorf("fail get users, %w", err)
		log.Printf("%v\n", err)
		return nil
	}

	// []*model.User[]*dataloader.Resultに変換する
	output := make([]*dataloader.Result, len(keys))
	for index, v := range users {
		output[index] = &dataloader.Result{Data: v, Error: nil}

	}
	return output
}

func LoadUser(ctx context.Context, userID string) (*model.User, error) {
	log.Printf("LoadUser(id = %s)\n", userID)
	loaders := GetLoaders(ctx)
	thunk := loaders.UserLoader.Load(ctx, dataloader.StringKey(userID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	user := result.(*model.User)
	log.Printf("return LoadUser(id = %d, name = %s)\n", user.ID, user.Name)
	return user, nil
}
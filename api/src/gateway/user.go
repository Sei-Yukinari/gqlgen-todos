package gateway

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/repository"
	"github.com/Sei-Yukinari/gqlgen-todos/src/util/apperror"
	"gorm.io/gorm"
)

type User struct {
	tx *gorm.DB
}

func NewUser(tx *gorm.DB) *User {
	return &User{
		tx: tx,
	}
}

var _ repository.User = (*User)(nil)

func (u User) FindByIDs(ctx context.Context, ids []int) ([]*model.User, apperror.AppError) {
	var users []*model.User
	if err := u.tx.
		Where("id IN ?", ids).
		Find(&users).
		Error; err != nil {
		return nil, apperror.Wrap(err).SetCode(apperror.Database)
	}
	return users, nil
}

package test

import (
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"gorm.io/gorm"
)

type Model interface {
	model.Todo | int
}

func Seeds(rdb *gorm.DB, seeds []interface{}) error {
	if rdb == nil {
		return nil
	}
	if seeds == nil {
		return nil
	}
	for _, s := range seeds {
		if err := rdb.Create(s).Error; err != nil {
			return err
		}
	}
	return nil
}

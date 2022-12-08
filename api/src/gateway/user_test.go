package gateway_test

import (
	"testing"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/Sei-Yukinari/gqlgen-todos/test"
	"github.com/stretchr/testify/assert"
)

func TestUser_FindByIDs(t *testing.T) {
	rdb := test.NewRDB(t, mysqlContainer)
	t.Run("GET Users By IDs", func(t *testing.T) {
		actual := []*model.User{
			{
				ID:   1,
				Name: "Dummy1",
			},
			{
				ID:   2,
				Name: "Dummy2",
			},
		}
		err := test.Seeds(rdb,
			[]interface{}{
				actual,
			})
		if err != nil {
			logger.Fatalf("fail seed data: %s", err)
		}
		repo := gateway.NewUser(rdb)
		res, err := repo.FindByIDs(ctx, []int{1, 2})
		assert.NoError(t, err)
		assert.Equal(t, res, actual)
	})
}

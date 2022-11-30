package gerror

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/Sei-Yukinari/gqlgen-todos/src/util/apperror"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func HandleError(ctx context.Context, err apperror.AppError) {

	logger.Warnf("Error : %s, Message : %s", err.Error(), err.InfoMessage())

	switch err.Code() {
	case apperror.Internal:
		AddErr(ctx, GetInfoMessage(err), err.Code())
	case apperror.Database:
		AddErr(ctx, GetInfoMessage(err), err.Code())
	default:
		AddErr(ctx, GetInfoMessage(err), err.Code())
	}
}

func AddErr(ctx context.Context, message string, code apperror.ErrorCode) {
	graphql.AddError(ctx, &gqlerror.Error{
		Message:    message,
		Path:       graphql.GetPath(ctx),
		Extensions: map[string]interface{}{"code": code},
	})
}

func GetInfoMessage(apperr apperror.AppError) string {
	if apperr.InfoMessage() != "" {
		return apperr.InfoMessage()
	}

	return "internal server error"
}

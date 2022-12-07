package test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"

	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
)

type GraphQLRequestBody struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables,omitempty"`
	OperationName string                 `json:"operationName,omitempty"`
}

func DoRequest(
	handler *gin.Engine,
	method string,
	target string,
	query string,
	variable map[string]interface{},
	contentType string,
) *httptest.ResponseRecorder {

	q := struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables,omitempty"`
	}{
		Query:     query,
		Variables: variable,
	}
	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		logger.Warnf("error encode%v", err)
	}
	r := httptest.NewRequest(method, target, &body)
	if contentType != "" {
		r.Header.Set("Content-Type", contentType)
	}
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)
	return w
}

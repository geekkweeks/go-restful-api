package middleware

import (
	"geekkweeks/go-restful-api/common"
	"geekkweeks/go-restful-api/helper"
	"geekkweeks/go-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

var authKey = "BwFvHRMEGQ"
var apiKey = "X-API-Key"

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Check the client the value of API key(X-API-Key) in the header is correct
	if authKey == request.Header.Get(apiKey) {
		// will continue to the next handler
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// error Unauthorize will be given
		writer.Header().Set(common.ContentType, common.ApplicationJson)
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}

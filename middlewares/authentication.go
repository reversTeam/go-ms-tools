package middlewares

import (
	"context"
	"net/http"

	"github.com/reversTeam/go-ms/core"
	"google.golang.org/grpc/metadata"
)

type AuthMiddleware struct {
	core.BaseMiddleware
}

func (a *AuthMiddleware) Apply(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, req, core.NewHttpError(http.StatusUnauthorized)
	}

	authHeader, exists := md["authorization"]
	if !exists || len(authHeader) == 0 {
		return ctx, req, core.NewHttpError(http.StatusUnauthorized, "authorization is missing")
	}

	expectedAuthValue := "Bearer toto"
	if authHeader[0] != expectedAuthValue {
		return ctx, req, core.NewHttpError(http.StatusUnauthorized, "Invalid token")
	}

	return ctx, req, nil
}

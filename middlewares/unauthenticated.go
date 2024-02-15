package middlewares

import (
	"context"

	"github.com/reversTeam/go-ms/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UnauthenticatedMiddleware struct {
	core.BaseMiddleware
}

func (a *UnauthenticatedMiddleware) Apply(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, req, status.Errorf(codes.Internal, "missing metadata")
	}

	authHeader, exists := md["authorization"]
	if !exists || len(authHeader) == 0 {
		return ctx, req, nil
	}

	expectedAuthValue := "Bearer toto"
	if authHeader[0] == expectedAuthValue {
		return ctx, req, status.Errorf(codes.Unauthenticated, "Unauthorized to perform this call with connexion")
	}

	return ctx, req, nil
}

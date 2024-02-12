package middlewares

import (
	"context"
	"fmt"

	"github.com/reversTeam/go-ms/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// func AuthMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 	fmt.Println("AUTH MIDDLEWARE IS APPLIED")

// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		return nil, status.Errorf(codes.Internal, "missing metadata")
// 	}

// 	authHeader, exists := md["authorization"]
// 	if !exists || len(authHeader) == 0 {
// 		return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
// 	}

// 	expectedAuthValue := "Bearer toto"
// 	if authHeader[0] != expectedAuthValue {
// 		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
// 	}

// 	return handler(ctx, req)
// }

type AuthMiddleware struct {
	core.BaseMiddleware
}

func (a *AuthMiddleware) Apply(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	fmt.Println("AUTH MIDDLEWARE IS APPLIED")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, req, status.Errorf(codes.Internal, "missing metadata")
	}

	authHeader, exists := md["authorization"]
	if !exists || len(authHeader) == 0 {
		return ctx, req, status.Errorf(codes.Unauthenticated, "missing authorization header")
	}

	expectedAuthValue := "Bearer toto"
	if authHeader[0] != expectedAuthValue {
		return ctx, req, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	return ctx, req, nil
}

package middlewares

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/reversTeam/go-ms/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckParametersMiddleware struct {
	core.BaseMiddleware
	validator *validator.Validate
}

func NewCheckParametersMiddleware() *CheckParametersMiddleware {
	return &CheckParametersMiddleware{
		validator: validator.New(),
	}
}

func (v *CheckParametersMiddleware) Apply(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	if v.validator == nil {
		return ctx, req, status.Errorf(codes.Internal, "Validator is not initialized")
	}

	err := v.validator.Struct(req)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errMsg string
			for _, err := range validationErrors {
				errMsg += fmt.Sprintf("Field validation for '%s' failed on the '%s' tag\n", err.Namespace(), err.Tag())
			}
			return ctx, req, status.Errorf(codes.InvalidArgument, "Validation failed: %s", errMsg)
		}

		return ctx, req, status.Errorf(codes.Internal, "Unexpected validation error: %v", err)
	}

	return ctx, req, nil
}

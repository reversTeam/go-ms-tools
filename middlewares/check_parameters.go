package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/reversTeam/go-ms/core"
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
		return ctx, req, core.NewHttpError(http.StatusInternalServerError)
	}

	err := v.validator.Struct(req)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errMsg string
			for _, err := range validationErrors {
				errMsg += fmt.Sprintf("Field validation for '%s' failed on the '%s' tag\n", err.Namespace(), err.Tag())
			}
			return ctx, req, core.NewHttpError(http.StatusBadRequest, "Validation failed: %s", errMsg)
		}

		return ctx, req, core.NewHttpError(http.StatusBadRequest, "Validation failed")
	}

	return ctx, req, nil
}

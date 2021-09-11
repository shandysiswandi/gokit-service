package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/shandysiswandi/gokit-service/pkg/jwt"
)

func JWTMiddleware(secret string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			if err := jwt.HasToken(ctx); err != nil {
				return nil, err
			}

			if isHasToken := jwt.ValidateToken(ctx); !isHasToken {
				return nil, jwt.ErrNoTokenFound
			}

			return next(ctx, request)
		}
	}
}

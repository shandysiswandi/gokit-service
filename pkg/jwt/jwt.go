package jwt

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

const (
	// Bearer is a prefix from value of header authorization
	Bearer string = "bearer"
)

type contextKey string

const (
	// JWTContextKey holds the key used to store a JWT in the context.
	JWTContextKey contextKey = "JWTToken"

	// JWTClaimsContextKey holds the key used to store the JWT Claims in the context.
	JWTClaimsContextKey contextKey = "JWTClaims"
)

var (
	// ErrNoTokenFound denote a token was not formatted as a JWT.
	ErrNoTokenFound = errors.New("no token found")
)

// HTTPToContext moves a JWT from request header to context. Particularly useful for servers.
func HTTPToContext(ctx context.Context, r *http.Request) context.Context {
	token, err := extractTokenFromAuthHeader(r.Header.Get("Authorization"))
	if err != nil {
		return ctx
	}
	return context.WithValue(ctx, JWTContextKey, token)
}

// HasToken to check if in context has a token
func HasToken(ctx context.Context) error {
	valCtx := ctx.Value(JWTContextKey)
	val, ok := valCtx.(string)
	if !ok || val == "" {
		return ErrNoTokenFound
	}
	return nil
}

// ValidateToken to check if in context has a token
// TODO: not implement
func ValidateToken(ctx context.Context) bool {
	valCtx := ctx.Value(JWTContextKey)
	val, ok := valCtx.(string)
	if !ok || val == "" {
		return false
	}
	return true
}

// extractTokenFromAuthHeader convert from header to value
func extractTokenFromAuthHeader(val string) (string, error) {
	part := strings.Split(val, " ")
	if len(part) != 2 || !strings.EqualFold(part[0], Bearer) {
		return "", ErrNoTokenFound
	}
	return part[1], nil
}

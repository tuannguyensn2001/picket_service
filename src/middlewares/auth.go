package middlewares

import (
	"context"
)

// exampleAuthFunc is used by a middleware to authenticate requests
func Auth(ctx context.Context) (context.Context, error) {

	// WARNING: in production define your own type to avoid context collisions
	newCtx := context.WithValue(ctx, "tokenInfo", "123")

	return newCtx, nil
}

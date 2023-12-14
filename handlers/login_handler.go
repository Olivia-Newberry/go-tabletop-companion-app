package handlers

import (
	"context"
)

func IsLoggedIn(ctx context.Context) bool {
	return ctx.Value("user") != nil
}

func GetLoggedInUser(ctx context.Context) string {
	return ctx.Value("user").(string)
}

func SetLoggedInUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, "user", user)
}

func ClearLoggedInUser(ctx context.Context) context.Context {
	return context.WithValue(ctx, "user", nil)
}

func Login(ctx context.Context) context.Context {
	return context.WithValue(ctx, "user", "test")
}

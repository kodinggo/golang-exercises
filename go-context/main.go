package main

import (
	"context"
	"fmt"
)

type ctxKey string

const usernameKey ctxKey = "username"
const useIDKey ctxKey = "userID"
const authKey ctxKey = "auth"

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, usernameKey, "john")
	ctx = context.WithValue(ctx, useIDKey, 1)
	ctx = context.WithValue(ctx, authKey, map[string]any{
		"username": "john",
		"userID":   1,
	})

	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	fmt.Println(ctx.Value(usernameKey))
	fmt.Println(ctx.Value(useIDKey))

	authData := ctx.Value(authKey)
	authDataMap, ok := authData.(map[string]any)
	if !ok {
		panic("auth data invalid")
	}
	username, ok := authDataMap["username"].(string)
	if ok {
		fmt.Println(username)
	}
}

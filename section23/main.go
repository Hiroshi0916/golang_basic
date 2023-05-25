package main

import (
	"context"
	"fmt"
)

type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func Set(userID, authToken string) context.Context {
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	return ctx
}

func Get(ctx context.Context) {
	fmt.Printf("userID: %v,authToken: %v\n", ctx.Value(ctxUserID), ctx.Value(ctxAuthToken))
}

func main() {
	ctx := Set("12345", "abc123")
	Get(ctx)
}

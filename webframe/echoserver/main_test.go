package main

import (
	"fmt"
	"golang.org/x/net/context"
	"testing"
)

func A(ctx context.Context) error {
	fmt.Println("A")
	return nil
}

func B(ctx context.Context) error {
	fmt.Println("B")
	return nil
}

func C(ctx context.Context) error {
	fmt.Println("C")
	return nil
}

func TestRecursive(t *testing.T) {
	var h func(context.Context) error

	h = A

	h = applyMiddleware(B, C)

	c := context.Background()

	h(c)
}

func applyMiddleware(h func(ctx context.Context) error, middleware ...func(ctx context.Context) error) func(ctx context.Context) error {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i]
	}
	return h
}

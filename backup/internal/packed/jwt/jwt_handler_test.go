package jwt_test

import (
	"context"
	"login-demo/internal/packed/jwt"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestGenerateToken(t *testing.T) {
	ctx := context.Background()
	gtest.C(t, func(t *gtest.T) {
		token, err := jwt.GenerateToken(ctx, "lusson")
		t.AssertNil(err)
		t.AssertNE(token, "")
	})
}

func TestValid(t *testing.T) {
	ctx := context.Background()
	gtest.C(t, func(t *gtest.T) {
		token, err := jwt.GenerateToken(ctx, "lusson")
		valid := jwt.Valid(ctx, token)
		t.AssertNil(err)
		t.Assert(valid, true)
	})
}

func TestParse(t *testing.T) {
	ctx := context.Background()
	gtest.C(t, func(t *gtest.T) {
		token, err := jwt.GenerateToken(ctx, "lusson")
		t.AssertNil(err)
		claims, err := jwt.Parse(ctx, token)
		t.AssertNil(err)
		t.Assert(claims.Username, "lusson")
	})
}

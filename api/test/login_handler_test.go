package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	ctx := RequestCtx("GET", "/api/v1/login", nil)

	assert.Equal(t, ctx.Response.StatusCode(), 200)
	assert.Equal(t, ctx.Response.Body(), []byte("Hi"))
}

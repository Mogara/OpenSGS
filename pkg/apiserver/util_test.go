package apiserver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllowOriginFunc(t *testing.T) {
	allowedOrigins := []string{"http://localhost:8000", "http://example.com", "http://test-*-app.example.com"}
	allowOriginFunc := allowOriginFunc(allowedOrigins)
	assert.Equal(t, true, allowOriginFunc("http://localhost:8000"))
	assert.Equal(t, false, allowOriginFunc("http://example.com:8000"))
	assert.Equal(t, true, allowOriginFunc("http://example.com"))
	assert.Equal(t, true, allowOriginFunc("http://test-hello-app.example.com"))
	assert.Equal(t, true, allowOriginFunc("http://test-world-app.example.com"))
	assert.Equal(t, false, allowOriginFunc("http://test1-hello-app.example.com"))
	assert.Equal(t, false, allowOriginFunc("http://test-hello-sss-app.example.com"))
}

func TestWildcardMatch(t *testing.T) {
	assert.Equal(t, true, wildcardMatch([]rune(""), []rune("")))
	assert.Equal(t, true, wildcardMatch([]rune("*"), []rune("123")))
	assert.Equal(t, true, wildcardMatch([]rune("12*45"), []rune("12345")))
	assert.Equal(t, true, wildcardMatch([]rune("12*45"), []rune("12333345")))
	assert.Equal(t, false, wildcardMatch([]rune("12*45"), []rune("123-3345")))
	assert.Equal(t, true, wildcardMatch([]rune("1-*-3"), []rune("1-2333-3")))
	assert.Equal(t, true, wildcardMatch([]rune("1-*-3"), []rune("1-3-3")))
	assert.Equal(t, false, wildcardMatch([]rune("1-*-3"), []rune("1-3-3-3")))
}

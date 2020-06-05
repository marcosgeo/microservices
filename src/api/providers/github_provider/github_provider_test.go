package providers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetuAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("u234ja03jaljfjw-9")
	assert.Equal(t, "token u234ja03jaljfjw-9", header)
}

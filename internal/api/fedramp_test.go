package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsFedramp(t *testing.T) {
	assert.True(t, IsFedramp("https://api.fedramp.vulnmapgov.io/something/else"))
	assert.True(t, IsFedramp("vulnmap.khulnasoft.com"))
	assert.False(t, IsFedramp("https://api.khulnasoft.com/"))
}

package internal

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHomePath(t *testing.T) {
	homedir := GetHomePath()
	assert.NotEqual(t, "", homedir)
	assert.True(t, strings.HasPrefix(homedir, "/"))
	assert.False(t, strings.HasSuffix(homedir, "/"))
}

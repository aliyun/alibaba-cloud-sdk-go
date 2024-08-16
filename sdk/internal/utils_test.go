package internal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDefaultString(t *testing.T) {
	assert.Equal(t, "default", GetDefaultString("", "default"))
	assert.Equal(t, "custom", GetDefaultString("custom", "default"))
	assert.Equal(t, "", GetDefaultString("", "", ""))
}

func TestMemoryAndRollback(t *testing.T) {
	os.Setenv("test", "old")
	rollback := Memory("test")
	os.Setenv("test", "new")
	rollback()

	assert.Equal(t, "old", os.Getenv("test"))
}

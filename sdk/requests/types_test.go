package requests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInteger(t *testing.T) {
	integer := NewInteger(123123)
	assert.True(t, integer.hasValue())
	value, err := integer.getValue()
	assert.Nil(t, err)
	assert.Equal(t, value, 123123)
}

func TestNewBoolean(t *testing.T) {
	boolean := NewBoolean(false)
	assert.True(t, boolean.hasValue())
	value, err := boolean.getValue()
	assert.Nil(t, err)
	assert.Equal(t, value, false)
}

func TestNewFloat(t *testing.T) {
	float := NewFloat(123123.123123)
	assert.True(t, float.hasValue())
	value, err := float.getValue()
	assert.Nil(t, err)
	assert.Equal(t, value, 123123.123123)
}

package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessKeyCredential(t *testing.T) {
	c := NewAccessKeyCredential("accesskeyid", "accesskeysecret")
	assert.Equal(t, "accesskeyid", c.AccessKeyId)
	assert.Equal(t, "accesskeysecret", c.AccessKeySecret)
}

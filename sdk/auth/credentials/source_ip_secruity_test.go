package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSourceCredential(t *testing.T) {
	c := NewSourceCredential("accesskeyid", "accesskeysecret", "sourceIp", "secruityTransport")
	assert.Equal(t, "accesskeyid", c.AccessKeyId)
	assert.Equal(t, "accesskeysecret", c.AccessKeySecret)
	assert.Equal(t, "sourceIp", c.SourceIp)
	assert.Equal(t, "secruityTransport", c.SecurityTransport)
}

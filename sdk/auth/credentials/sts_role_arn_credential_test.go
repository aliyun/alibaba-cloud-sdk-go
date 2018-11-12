package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleArnCredential(t *testing.T) {
	c := NewRamRoleArnCredential("accessKeyId", "accessKeySecret", "roleArn", "roleSessionName", 3600)
	assert.Equal(t, "accessKeyId", c.AccessKeyId)
	assert.Equal(t, "accessKeySecret", c.AccessKeySecret)
	assert.Equal(t, "roleArn", c.RoleArn)
	assert.Equal(t, "roleSessionName", c.RoleSessionName)
	assert.Equal(t, 3600, c.RoleSessionExpiration)
}

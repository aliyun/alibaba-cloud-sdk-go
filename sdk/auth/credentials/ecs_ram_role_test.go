package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestECSRamRole(t *testing.T) {
	c := NewEcsRamRoleCredential("rolename")
	assert.Equal(t, "rolename", c.RoleName)
}

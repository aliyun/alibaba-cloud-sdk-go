package endpoints

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMappingResolver_TryResolve(t *testing.T) {

	regionId := "cn-hangzhou"
	productId := "ecs"
	endpoint := GetEndpointFromMap(regionId, productId)
	assert.Equal(t, "", endpoint)

	AddEndpointMapping("cn-hangzhou", "Ecs", "unreachable.aliyuncs.com")

	endpoint = GetEndpointFromMap(regionId, productId)
	assert.Equal(t, "unreachable.aliyuncs.com", endpoint)
}

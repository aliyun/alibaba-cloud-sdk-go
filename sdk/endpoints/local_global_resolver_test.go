package endpoints

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestLocalGlobalResolver_GetName(t *testing.T) {
  resolver := &LocalGlobalResolver{}
  assert.Equal(t, resolver.GetName(), "local global resolver")
}

func TestLocalGlobalResolver_TryResolve(t *testing.T) {
  resolver := &LocalGlobalResolver{}
  resolveParam := &ResolveParam{
		Product: "ecs",
	}
  endpoint, support, err := resolver.TryResolve(resolveParam)
  assert.Nil(t, err)
  assert.Equal(t, "", endpoint)
	assert.Equal(t, false, support)

	resolveParam = &ResolveParam{
		Product: "drds",
	}
  endpoint, support, err = resolver.TryResolve(resolveParam)
  assert.Nil(t, err)
  assert.Equal(t, "drds.aliyuncs.com", endpoint)
	assert.Equal(t, true, support)

	resolveParam = &ResolveParam{
		Product: "inexist",
	}
  endpoint, support, err = resolver.TryResolve(resolveParam)
  assert.Nil(t, err)
  assert.Equal(t, "", endpoint)
	assert.Equal(t, false, support)
}

// func TestLocationResolver_TryResolve_LocationWithError(t *testing.T) {
//   resolver := &LocationResolver{}
//   resolveParam := &ResolveParam{
//     LocationProduct: "ecs",
//     RegionId: "cn-hangzhou",
//     Product:  "ecs",
//     CommonApi: func(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
//       err = errors.NewClientError("SDK.MockError", "Mock error", nil)
//       return
//     },
//   }
//   endpoint, support, err := resolver.TryResolve(resolveParam)
//   assert.Equal(t, "", endpoint)
//   assert.Equal(t, false, support)
//   assert.Equal(t, "[SDK.MockError] Mock error", err.Error())
// }

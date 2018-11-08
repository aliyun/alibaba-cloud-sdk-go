package endpoints

import (
  "testing"

  "github.com/stretchr/testify/assert"

  "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
  "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
  "github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func TestLocationResolver_GetName(t *testing.T) {
  resolver := &LocationResolver{}
  assert.Equal(t, "location resolver", resolver.GetName())
}

func TestLocationResolver_TryResolve_EmptyLocationProduct(t *testing.T) {
  resolver := &LocationResolver{}
  resolveParam := &ResolveParam{}
  endpoint, support, err := resolver.TryResolve(resolveParam)
  assert.Nil(t, err)
  assert.Equal(t, "", endpoint)
  assert.Equal(t, false, support)
}

func TestLocationResolver_TryResolve_LocationWithError(t *testing.T) {
  resolver := &LocationResolver{}
  resolveParam := &ResolveParam{
    LocationProduct: "ecs",
    RegionId: "cn-hangzhou",
    Product:  "ecs",
    CommonApi: func(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
      err = errors.NewClientError("SDK.MockError", "Mock error", nil)
      return
    },
  }
  endpoint, support, err := resolver.TryResolve(resolveParam)
  assert.Equal(t, "", endpoint)
  assert.Equal(t, false, support)
  assert.Equal(t, "[SDK.MockError] Mock error", err.Error())
}

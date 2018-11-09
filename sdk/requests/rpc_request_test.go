package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RpcRequest(t *testing.T) {
	r := &RpcRequest{}
	r.InitWithApiInfo("product", "version", "action", "serviceCode", "endpointType")
	assert.NotNil(t, r)

	assert.Equal(t, "POST", r.GetMethod())
	assert.Equal(t, "RPC", r.GetStyle())
	assert.Equal(t, "product", r.GetProduct())
	assert.Equal(t, "version", r.GetVersion())
	assert.Equal(t, "action", r.GetActionName())
	assert.Equal(t, "serviceCode", r.GetLocationServiceCode())
	assert.Equal(t, "endpointType", r.GetLocationEndpointType())

	// url
}

package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCommonRequest(t *testing.T) {
	r := NewCommonRequest()
	assert.NotNil(t, r)

	assert.Equal(t, "common", r.GetHeaders()["x-sdk-invoke-type"])
	assert.Equal(t, 0, len(r.PathParams))

	r.addPathParam("name", "value")
	assert.Equal(t, "value", r.PathParams["name"])
}

func Test_CommonRequest_TransToAcsRequest(t *testing.T) {
	r := NewCommonRequest()
	assert.NotNil(t, r)
	r.TransToAcsRequest()

	assert.Equal(t, "RPC", r.GetStyle())
	assert.Equal(t, "", r.GetQueries())
}

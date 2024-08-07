package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func mockRpcDate(fn func() string) string {
	return "mock date"
}

func mockRpcGetNonce(fn func() string) string {
	return "MOCK_UUID"
}

func TestRpcSignatureComposer_buildRpcStringToSign(t *testing.T) {
	request := requests.NewCommonRequest()
	request.TransToAcsRequest()
	stringToSign := buildRpcStringToSign(request)
	assert.Equal(t, "GET&%2F&", stringToSign)
	request.FormParams["key"] = "value"
	stringToSign = buildRpcStringToSign(request)
	assert.Equal(t, "GET&%2F&key%3Dvalue", stringToSign)
	request.QueryParams["q"] = "value"
	stringToSign = buildRpcStringToSign(request)
	assert.Equal(t, "GET&%2F&key%3Dvalue%26q%3Dvalue", stringToSign)
	request.QueryParams["q"] = "http://domain/?q=value&q2=value2"
	stringToSign = buildRpcStringToSign(request)
	assert.Equal(t, "GET&%2F&key%3Dvalue%26q%3Dhttp%253A%252F%252Fdomain%252F%253Fq%253Dvalue%2526q2%253Dvalue2", stringToSign)
}

func TestRpcSignatureComposer(t *testing.T) {
	request := requests.NewCommonRequest()
	request.TransToAcsRequest()
	c := credentials.NewAccessKeyCredential("accessKeyId", "accessKeySecret")

	origTestHookGetDate := hookGetDate
	defer func() { hookGetDate = origTestHookGetDate }()
	hookGetDate = mockRpcDate
	origTestHookGetNonce := hookGetNonce
	defer func() { hookGetNonce = origTestHookGetNonce }()
	hookGetNonce = mockRpcGetNonce

	provider, err := ToCredentialsProvider(c)
	assert.Nil(t, err)
	signRpcRequest(request, "regionId", provider)
	assert.Equal(t, "mock date", request.GetQueryParams()["Timestamp"])
	assert.Equal(t, "MOCK_UUID", request.GetQueryParams()["SignatureNonce"])
	assert.Equal(t, "7loPmFjvDnzOVnQeQNj85S6nFGY=", request.GetQueryParams()["Signature"])

	sts := credentials.NewStsTokenCredential("accessKeyId", "accessKeySecret", "accessKeyStsToken")

	provider, err = ToCredentialsProvider(sts)
	assert.Nil(t, err)
	signRpcRequest(request, "regionId", provider)
	assert.Equal(t, "mock date", request.GetQueryParams()["Timestamp"])
	assert.Equal(t, "MOCK_UUID", request.GetQueryParams()["SignatureNonce"])
	assert.Equal(t, "5Nxdcler+ihqWqv0Hr2On4PsBf4=", request.GetQueryParams()["Signature"])
}

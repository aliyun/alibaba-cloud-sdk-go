package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func mockDate(fn func() string) string {
	return "mock date"
}

func TestRoaSignatureComposer_buildRoaStringToSign(t *testing.T) {
	request := requests.NewCommonRequest()
	request.PathPattern = "/users/:user"
	request.TransToAcsRequest()
	stringToSign := buildRoaStringToSign(request)
	assert.Equal(t, "GET\nx-acs-version:\n/users/:user", stringToSign)

	request.Headers["Accept"] = "application/json;charset=utf8"
	stringToSign = buildRoaStringToSign(request)
	assert.Equal(t, "GET\napplication/json;charset=utf8\nx-acs-version:\n/users/:user", stringToSign)
	request.SetContentType("application/json")
	request.Headers["x-acs-custom-header"] = "value"
	stringToSign = buildRoaStringToSign(request)
	assert.Equal(t, "GET\napplication/json;charset=utf8\napplication/json\nx-acs-custom-header:value\nx-acs-version:\n/users/:user", stringToSign)
	request.QueryParams["q"] = "value"
	stringToSign = buildRoaStringToSign(request)
	assert.Equal(t, "GET\napplication/json;charset=utf8\napplication/json\nx-acs-custom-header:value\nx-acs-version:\n/users/:user?q=value", stringToSign)
	request.QueryParams["q"] = "http://domain/?q=value&q2=value2"
	stringToSign = buildRoaStringToSign(request)
	assert.Equal(t, "GET\napplication/json;charset=utf8\napplication/json\nx-acs-custom-header:value\nx-acs-version:\n/users/:user?q=http://domain/?q=value&q2=value2", stringToSign)
}

func TestRoaSignatureComposer(t *testing.T) {
	request := requests.NewCommonRequest()
	request.PathPattern = "/users/:user"
	request.TransToAcsRequest()
	c := credentials.NewAccessKeyCredential("accessKeyId", "accessKeySecret")

	origTestHookGetDate := hookGetDate
	defer func() { hookGetDate = origTestHookGetDate }()
	hookGetDate = mockDate

	provider, err := ToCredentialsProvider(c)
	assert.Nil(t, err)
	signRoaRequest(request, provider)
	assert.Equal(t, "mock date", request.GetHeaders()["Date"])
	assert.Equal(t, "static_ak", request.GetHeaders()["x-acs-credentials-provider"])
	assert.Equal(t, "acs accessKeyId:PPPSwv0LPoLvg6rGxp7dd4uPESo=", request.GetHeaders()["Authorization"])
}

func TestRoaSignatureComposer2(t *testing.T) {
	request := requests.NewCommonRequest()
	request.PathPattern = "/users/:user"
	request.FormParams["key"] = "value"
	request.AcceptFormat = "XML"
	request.TransToAcsRequest()
	c := credentials.NewAccessKeyCredential("accessKeyId", "accessKeySecret")

	origTestHookLookupIP := hookGetDate
	defer func() { hookGetDate = origTestHookLookupIP }()
	hookGetDate = mockDate

	provider, err := ToCredentialsProvider(c)
	assert.Nil(t, err)
	signRoaRequest(request, provider)
	assert.Equal(t, "application/x-www-form-urlencoded", request.GetHeaders()["Content-Type"])
	assert.Equal(t, "mock date", request.GetHeaders()["Date"])
	assert.Equal(t, "application/xml", request.GetHeaders()["Accept"])
	assert.Equal(t, "static_ak", request.GetHeaders()["x-acs-credentials-provider"])
	assert.Equal(t, "acs accessKeyId:X/tjQCOP4ZMCzQ1I1IdCra3byPA=", request.GetHeaders()["Authorization"])
}

func TestRoaSignatureComposer3(t *testing.T) {
	request := requests.NewCommonRequest()
	request.PathPattern = "/users/:user"
	request.AcceptFormat = "RAW"
	request.TransToAcsRequest()
	c := credentials.NewAccessKeyCredential("accessKeyId", "accessKeySecret")

	origTestHookGetDate := hookGetDate
	defer func() { hookGetDate = origTestHookGetDate }()
	hookGetDate = mockDate

	provider, err := ToCredentialsProvider(c)
	assert.Nil(t, err)
	signRoaRequest(request, provider)
	assert.Equal(t, "mock date", request.GetHeaders()["Date"])
}

func TestCompleteROASignParams(t *testing.T) {
	req := requests.NewCommonRequest()
	req.TransToAcsRequest()
	c := credentials.NewBearerTokenCredential("Bearer.Token")
	provider, err := ToCredentialsProvider(c)
	assert.Nil(t, err)
	cc, err := provider.GetCredentials()
	assert.Nil(t, err)

	completeROASignParams(req, cc)
	head := req.GetHeaders()
	assert.Equal(t, "Bearer.Token", head["x-acs-bearer-token"])

	stc := credentials.NewStsTokenCredential("accessKeyId", "accessKeySecret", "accessKeyStsToken")
	provider, err = ToCredentialsProvider(stc)
	assert.Nil(t, err)
	cc, err = provider.GetCredentials()
	assert.Nil(t, err)
	completeROASignParams(req, cc)
	head = req.GetHeaders()
	assert.Equal(t, "accessKeyStsToken", head["x-acs-security-token"])
}

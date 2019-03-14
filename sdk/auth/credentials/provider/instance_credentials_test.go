package provider_test

import (
	"os"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"

	"github.com/stretchr/testify/assert"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials/provider"
)

func TestInstanceCredentialsProvider(t *testing.T) {
	p := provider.NewInstanceCredentialsProvider()
	c, err := p.Resolve()
	assert.Nil(t, c)
	assert.Nil(t, err)

	os.Setenv(provider.ENVEcsMetadata, "")
	c, err = p.Resolve()
	assert.Nil(t, c)
	assert.EqualError(t, err, "Environmental variable 'ALIBABA_CLOUD_ECS_METADATA' are empty")

	os.Setenv(provider.ENVEcsMetadata, "test")
	c, err = p.Resolve()
	assert.Nil(t, c)
	assert.NotNil(t, err)

	provider.HookGet = func(fn func(string) (int, []byte, error)) func(string) (int, []byte, error) {
		return func(string) (int, []byte, error) {
			return 404, []byte(""), nil
		}
	}
	c, err = p.Resolve()
	assert.Nil(t, c)
	assert.EqualError(t, err, "The role was not found in the instance")

	provider.HookGet = func(fn func(string) (int, []byte, error)) func(string) (int, []byte, error) {
		return func(string) (int, []byte, error) {
			return 400, []byte(""), nil
		}
	}
	c, err = p.Resolve()
	assert.Nil(t, c)
	assert.EqualError(t, err, "Received 400 when getting security credentials for test")

	provider.HookGet = func(fn func(string) (int, []byte, error)) func(string) (int, []byte, error) {
		return func(string) (int, []byte, error) {
			return 200, []byte(`{
				"AccessKeyId" : "STS.*******",
				"AccessKeySecret" : "*******",
				"Expiration" : "2019-01-28T15:15:56Z",
				"SecurityToken" : "bbbbb",
				"LastUpdated" : "2019-01-28T09:15:55Z",
				"Code" : "Success"
			  }`), nil
		}
	}
	c, err = p.Resolve()
	assert.Nil(t, err)
	assert.Equal(t, credentials.NewStsTokenCredential("STS.*******", "*******", "bbbbb"), c)

}

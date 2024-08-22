package credentials

import (
	"os"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/internal"
	"github.com/stretchr/testify/assert"
)

func TestEnvironmentVariableCredentialsProvider(t *testing.T) {
	rollback := internal.Memory("ALIBABA_CLOUD_ACCESS_KEY_ID", "ALIBABA_CLOUD_ACCESS_KEY_SECRET", "ALIBABA_CLOUD_SECURITY_TOKEN")
	defer rollback()

	p := NewEnvironmentVariableCredentialsProvider()
	assert.NotNil(t, p)
	_, err := p.GetCredentials()
	assert.EqualError(t, err, "unable to get credentials from enviroment variables, Access key ID must be specified via environment variable (ALIBABA_CLOUD_ACCESS_KEY_ID)")
	os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_ID", "akid")
	_, err = p.GetCredentials()
	assert.EqualError(t, err, "unable to get credentials from enviroment variables, Access key secret must be specified via environment variable (ALIBABA_CLOUD_ACCESS_KEY_SECRET)")

	os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET", "aksecret")
	cc, err := p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid", cc.AccessKeyId)
	assert.Equal(t, "aksecret", cc.AccessKeySecret)
	assert.Equal(t, "", cc.SecurityToken)
	assert.Equal(t, "", cc.BearerToken)

	os.Setenv("ALIBABA_CLOUD_SECURITY_TOKEN", "token")
	cc, err = p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid", cc.AccessKeyId)
	assert.Equal(t, "aksecret", cc.AccessKeySecret)
	assert.Equal(t, "token", cc.SecurityToken)
	assert.Equal(t, "", cc.BearerToken)

	assert.Equal(t, "env", cc.ProviderName)
}

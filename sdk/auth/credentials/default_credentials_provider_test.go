package credentials

import (
	"errors"
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/internal"
	"github.com/stretchr/testify/assert"
)

func TestDefaultCredentialsProvider(t *testing.T) {
	provider := NewDefaultCredentialsProvider()
	assert.NotNil(t, provider)
	assert.Len(t, provider.providerChain, 4)
	_, ok := provider.providerChain[0].(*EnvironmentVariableCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[1].(*CLIProfileCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[2].(*ProfileCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[3].(*ECSRAMRoleCredentialsProvider)
	assert.True(t, ok)

	// Add oidc provider
	rollback := internal.Memory("ALIBABA_CLOUD_OIDC_TOKEN_FILE",
		"ALIBABA_CLOUD_OIDC_PROVIDER_ARN",
		"ALIBABA_CLOUD_ROLE_ARN",
		"ALIBABA_CLOUD_ECS_METADATA",
		"ALIBABA_CLOUD_CREDENTIALS_URI")

	defer rollback()
	os.Setenv("ALIBABA_CLOUD_OIDC_TOKEN_FILE", "/path/to/oidc.token")
	os.Setenv("ALIBABA_CLOUD_OIDC_PROVIDER_ARN", "oidcproviderarn")
	os.Setenv("ALIBABA_CLOUD_ROLE_ARN", "rolearn")

	provider = NewDefaultCredentialsProvider()
	assert.NotNil(t, provider)
	assert.Len(t, provider.providerChain, 5)
	_, ok = provider.providerChain[0].(*EnvironmentVariableCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[1].(*OIDCCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[2].(*CLIProfileCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[3].(*ProfileCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[4].(*ECSRAMRoleCredentialsProvider)
	assert.True(t, ok)

	// Add ecs ram role
	os.Setenv("ALIBABA_CLOUD_ECS_METADATA", "rolename")
	provider = NewDefaultCredentialsProvider()
	assert.NotNil(t, provider)
	assert.Len(t, provider.providerChain, 5)
	_, ok = provider.providerChain[0].(*EnvironmentVariableCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[1].(*OIDCCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[2].(*CLIProfileCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[3].(*ProfileCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[4].(*ECSRAMRoleCredentialsProvider)
	assert.True(t, ok)

	os.Setenv("ALIBABA_CLOUD_CREDENTIALS_URI", "http://")
	provider = NewDefaultCredentialsProvider()
	assert.NotNil(t, provider)
	assert.Len(t, provider.providerChain, 6)
	_, ok = provider.providerChain[0].(*EnvironmentVariableCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[1].(*OIDCCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[2].(*CLIProfileCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[3].(*ProfileCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[4].(*ECSRAMRoleCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[5].(*URLCredentialsProvider)
	assert.True(t, ok)
}

func TestDefaultCredentialsProvider_GetCredentials(t *testing.T) {
	rollback := internal.Memory("ALIBABA_CLOUD_ACCESS_KEY_ID",
		"ALIBABA_CLOUD_ACCESS_KEY_SECRET",
		"ALIBABA_CLOUD_SECURITY_TOKEN",
		"ALIBABA_CLOUD_ECS_METADATA_DISABLED",
		"ALIBABA_CLOUD_PROFILE")

	defer func() {
		getHomePath = internal.GetHomePath
		rollback()
	}()
	originDo := hookDo
	defer func() { hookDo = originDo }()

	// testcase: empty home
	getHomePath = func() string {
		return ""
	}

	os.Setenv("ALIBABA_CLOUD_ECS_METADATA_DISABLED", "true")
	provider := NewDefaultCredentialsProvider()
	assert.Len(t, provider.providerChain, 3)
	_, err := provider.GetCredentials()
	assert.EqualError(t, err, "unable to get credentials from any of the providers in the chain: unable to get credentials from enviroment variables, Access key ID must be specified via environment variable (ALIBABA_CLOUD_ACCESS_KEY_ID), cannot found home dir, cannot found home dir")

	os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_ID", "akid")
	os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET", "aksecret")
	provider = NewDefaultCredentialsProvider()
	assert.Len(t, provider.providerChain, 3)
	cc, err := provider.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, &Credentials{
		AccessKeyId:     "akid",
		AccessKeySecret: "aksecret",
		SecurityToken:   "",
		BearerToken:     "",
		ProviderName:    "default/env",
	}, cc)
	// get again
	cc, err = provider.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, &Credentials{
		AccessKeyId:     "akid",
		AccessKeySecret: "aksecret",
		SecurityToken:   "",
		BearerToken:     "",
		ProviderName:    "default/env",
	}, cc)

	getHomePath = func() string {
		wd, _ := os.Getwd()
		return path.Join(wd, "fixtures")
	}
	os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_ID", "")
	os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET", "")
	os.Setenv("ALIBABA_CLOUD_PROFILE", "ChainableRamRoleArn")
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"Credentials": {"AccessKeyId":"akid","AccessKeySecret":"aksecret","Expiration":"2021-10-20T04:27:09Z","SecurityToken":"ststoken"}}`)
			return
		}
	}
	provider = NewDefaultCredentialsProvider()
	cc, err = provider.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, &Credentials{AccessKeyId: "akid", AccessKeySecret: "aksecret", SecurityToken: "ststoken", ProviderName: "default/cli_profile/ram_role_arn/ram_role_arn/static_ak"}, cc)

	provider.lastUsedProvider = new(testProvider)
	cc, err = provider.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "test", cc.AccessKeyId)
	assert.Equal(t, "test", cc.AccessKeySecret)
	assert.Equal(t, "default/test", cc.ProviderName)

	provider.lastUsedProvider = new(testErrorProvider)
	_, err = provider.GetCredentials()
	assert.Equal(t, "error", err.Error())
}

type testProvider struct {
}

func (provider *testProvider) GetCredentials() (cc *Credentials, err error) {
	cc = &Credentials{
		AccessKeyId:     "test",
		AccessKeySecret: "test",
		ProviderName:    "",
	}
	return
}

func (provider *testProvider) GetProviderName() string {
	return "test"
}

type testErrorProvider struct {
}

func (provider *testErrorProvider) GetCredentials() (cc *Credentials, err error) {
	err = errors.New("error")
	return
}

func (provider *testErrorProvider) GetProviderName() string {
	return "test"
}

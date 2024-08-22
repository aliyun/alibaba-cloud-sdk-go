package credentials

import (
	"os"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/internal"
	"github.com/stretchr/testify/assert"
)

func TestDefaultCredentialsProvider(t *testing.T) {
	provider := NewDefaultCredentialsProvider()
	assert.NotNil(t, provider)
	assert.Len(t, provider.providerChain, 3)
	_, ok := provider.providerChain[0].(*EnvironmentVariableCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[1].(*CLIProfileCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[2].(*ProfileCredentialsProvider)
	assert.True(t, ok)

	// Add oidc provider
	rollback := internal.Memory("ALIBABA_CLOUD_OIDC_TOKEN_FILE",
		"ALIBABA_CLOUD_OIDC_PROVIDER_ARN",
		"ALIBABA_CLOUD_ROLE_ARN",
		"ALIBABA_CLOUD_ECS_METADATA")

	defer rollback()
	os.Setenv("ALIBABA_CLOUD_OIDC_TOKEN_FILE", "/path/to/oidc.token")
	os.Setenv("ALIBABA_CLOUD_OIDC_PROVIDER_ARN", "oidcproviderarn")
	os.Setenv("ALIBABA_CLOUD_ROLE_ARN", "rolearn")

	provider = NewDefaultCredentialsProvider()
	assert.NotNil(t, provider)
	assert.Len(t, provider.providerChain, 4)
	_, ok = provider.providerChain[0].(*EnvironmentVariableCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[1].(*OIDCCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[2].(*CLIProfileCredentialsProvider)
	assert.True(t, ok)

	_, ok = provider.providerChain[3].(*ProfileCredentialsProvider)
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
}

func TestDefaultCredentialsProvider_GetCredentials(t *testing.T) {
	rollback := internal.Memory("ALIBABA_CLOUD_ACCESS_KEY_ID",
		"ALIBABA_CLOUD_ACCESS_KEY_SECRET",
		"ALIBABA_CLOUD_SECURITY_TOKEN")

	defer func() {
		getHomePath = internal.GetHomePath
		rollback()
	}()

	// testcase: empty home
	getHomePath = func() string {
		return ""
	}

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
}

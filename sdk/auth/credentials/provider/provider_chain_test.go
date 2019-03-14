package provider_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials/provider"
)

func TestProviderChain(t *testing.T) {
	env := provider.NewEnvProvider()
	pp := provider.NewProfileProvider()
	instanceP := provider.NewInstanceCredentialsProvider()

	pc := provider.NewProviderChain([]provider.Provider{env, pp, instanceP})

	c, err := pc.Resolve()
	assert.Equal(t, &credentials.AccessKeyCredential{AccessKeyId: "AccessKeyId", AccessKeySecret: "AccessKeySecret"}, c)
	assert.Nil(t, err)
}

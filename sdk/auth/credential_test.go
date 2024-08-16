package auth

import (
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/stretchr/testify/assert"
)

func TestToCredentialsProvider(t *testing.T) {
	// StaticAKCredentialsProvider
	ak := credentials.NewAccessKeyCredential("id", "secret")
	p, err := ToCredentialsProvider(ak)
	assert.Nil(t, err)
	_, ok := p.(*credentials.StaticAKCredentialsProvider)
	assert.True(t, ok)

	// StaticAKCredentialsProvider
	sts := credentials.NewStsTokenCredential("id", "secret", "token")
	p, err = ToCredentialsProvider(sts)
	assert.Nil(t, err)
	_, ok = p.(*credentials.StaticSTSCredentialsProvider)
	assert.True(t, ok)

	bt := credentials.NewBearerTokenCredential("bearertoken")
	p, err = ToCredentialsProvider(bt)
	assert.Nil(t, err)
	_, ok = p.(*credentials.BearerTokenCredentialsProvider)
	assert.True(t, ok)

	errc := credentials.NewEcsRamRoleCredential("rolename")
	p, err = ToCredentialsProvider(errc)
	assert.Nil(t, err)
	_, ok = p.(*credentials.ECSRAMRoleCredentialsProvider)
	assert.True(t, ok)

	rpc := credentials.NewRsaKeyPairCredential("privatekey", "publickeyId", 900)
	p, err = ToCredentialsProvider(rpc)
	assert.Nil(t, err)
	_, ok = p.(*credentials.RSAKeyPairCredentialsProvider)
	assert.True(t, ok)

	rrac := credentials.NewRamRoleArnCredential("akid", "aksecret", "roleArn", "rsn", 900)
	p, err = ToCredentialsProvider(rrac)
	assert.Nil(t, err)
	_, ok = p.(*credentials.RAMRoleARNCredentialsProvider)
	assert.True(t, ok)

	oidc, err := credentials.NewOIDCCredentialsProviderBuilder().
		WithOIDCTokenFilePath("/path/to/oidctoken").
		WithOIDCProviderARN("providerarn").
		WithRoleArn("rolearn").
		Build()
	assert.Nil(t, err)
	p, err = ToCredentialsProvider(oidc)
	assert.Nil(t, err)
	_, ok = p.(*credentials.OIDCCredentialsProvider)
	assert.True(t, ok)

	// Default credentials provider
	p, err = ToCredentialsProvider(nil)
	assert.Nil(t, err)
	_, ok = p.(*credentials.DefaultCredentialsProvider)
	assert.True(t, ok)

	// unsupported
	_, err = ToCredentialsProvider("string")
	assert.NotNil(t, err)
	assert.Equal(t, "[SDK.UnsupportedCredential] Specified credential (type = string) is not supported, please check", err.Error())
}

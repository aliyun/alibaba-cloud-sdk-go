package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticAKCredentialsProvider(t *testing.T) {
	p := NewStaticAKCredentialsProvider("akid", "aksecret")
	c, err := p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid", c.AccessKeyId)
	assert.Equal(t, "aksecret", c.AccessKeySecret)
	assert.Equal(t, "", c.SecurityToken)
	assert.Equal(t, "", c.BearerToken)
}

func TestStaticSTSCredentialsProvider(t *testing.T) {
	p := NewStaticSTSCredentialsProvider("akid", "aksecret", "token")
	c, err := p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid", c.AccessKeyId)
	assert.Equal(t, "aksecret", c.AccessKeySecret)
	assert.Equal(t, "token", c.SecurityToken)
	assert.Equal(t, "", c.BearerToken)
}

func TestBearerTokenCredentialsProvider(t *testing.T) {
	p := NewBearerTokenCredentialsProvider("bearertoken")
	c, err := p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "", c.AccessKeyId)
	assert.Equal(t, "", c.AccessKeySecret)
	assert.Equal(t, "", c.SecurityToken)
	assert.Equal(t, "bearertoken", c.BearerToken)
}

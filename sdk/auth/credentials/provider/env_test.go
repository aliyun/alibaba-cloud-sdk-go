package provider

import (
	"os"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"

	"github.com/stretchr/testify/assert"
)

func TestEnvResolve(t *testing.T) {
	p := NewEnvProvider()
	assert.Equal(t, &EnvProvider{}, p)
	c, err := p.Resolve()
	assert.Nil(t, c)
	assert.Nil(t, err)
	os.Setenv(ENVAccessKeyID, "")
	os.Setenv(ENVAccessKeySecret, "")
	c, err = p.Resolve()
	assert.Nil(t, c)
	assert.EqualError(t, err, "Environmental variable (ALIBABACLOUD_ACCESS_KEY_ID or ALIBABACLOUD_ACCESS_KEY_SECRET) is empty")
	os.Setenv(ENVAccessKeyID, "AccessKeyId")
	os.Setenv(ENVAccessKeySecret, "AccessKeySecret")
	c, err = p.Resolve()
	assert.Nil(t, err)
	assert.Equal(t, &credentials.AccessKeyCredential{AccessKeyId: "AccessKeyId", AccessKeySecret: "AccessKeySecret"}, c)
}

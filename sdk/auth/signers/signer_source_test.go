package signers

import (
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/stretchr/testify/assert"
)

func TestNewSourceSigner(t *testing.T) {
	c := credentials.NewSourceCredential("accessKeyId", "accessKeySecret", "sourceIp", "secruityTransport")
	assert.NotNil(t, c)
	s := NewSourceSigner(c)
	assert.Nil(t, s.GetExtraParam())
	assert.Equal(t, "HMAC-SHA1", s.GetName())
	assert.Equal(t, "", s.GetType())
	assert.Equal(t, "1.0", s.GetVersion())
	accessKeyId, err := s.GetAccessKeyId()
	assert.Nil(t, err)
	assert.Equal(t, "accessKeyId", accessKeyId)
	assert.Equal(t, "Dqy7QZhP4TyQUDa3SBSFXopJaIo=", s.Sign("string to sign", "suffix"))
}

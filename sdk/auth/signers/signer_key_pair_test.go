package signers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
)

func TestKeyPairError(t *testing.T) {
	c := credentials.NewRsaKeyPairCredential("privateKey", "publicKey", 1)
	_, err := NewSignerKeyPair(c, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "[SDK.InvalidParam] Key Pair session duration should be in the range of 15min - 1Hr", err.Error())
}

func TestKeyPairOk(t *testing.T) {
	c := credentials.NewRsaKeyPairCredential("privateKey", "publicKey", 0)
	s, err := NewSignerKeyPair(c, nil)
	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, 3600, s.credentialExpiration)
	c = credentials.NewRsaKeyPairCredential("privateKey", "publicKey", 3500)
	s, err = NewSignerKeyPair(c, nil)
	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, 3500, s.credentialExpiration)
	assert.Equal(t, "HMAC-SHA1", s.GetName())
	assert.Equal(t, "1.0", s.GetVersion())
	assert.Equal(t, "", s.GetType())
	// nothing
	s.Shutdown()
}

// func TestGetAccessKeyId(t *testing.T) {
// 	c := credentials.NewRsaKeyPairCredential("privateKey", "publicKey", 0)
// 	s, err := NewSignerKeyPair(c, nil)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, s)
// 	accessKeyId, err := s.GetAccessKeyId()
// 	assert.NotNil(t, accessKeyId)
// }

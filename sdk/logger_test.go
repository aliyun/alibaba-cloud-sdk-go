package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OpenLogger(t *testing.T) {
	client, err := NewClientWithAccessKey("regionid", "acesskeyid", "accesskeysecret")
	assert.Nil(t, err)
	assert.NotNil(t, client)

	client.OpenLogger()
	assert.Equal(t, true, client.logger.isOpen)
}

func Test_SetTemplate(t *testing.T) {
	client, err := NewClientWithAccessKey("regionid", "acesskeyid", "accesskeysecret")
	assert.Nil(t, err)
	assert.NotNil(t, client)

	template := "{time}"
	client.SetTemplate(template)
	assert.Equal(t, "{time}", client.logger.formatTemplate)
}

func Test_GetTemplate(t *testing.T) {
	client, err := NewClientWithAccessKey("regionid", "acesskeyid", "accesskeysecret")
	assert.Nil(t, err)
	assert.NotNil(t, client)

	assert.Equal(t, defaultLoggerTemplate, client.GetTemplate())
}

func Test_GetLoggerMsg(t *testing.T) {
	client, err := NewClientWithAccessKey("regionid", "acesskeyid", "accesskeysecret")
	assert.Nil(t, err)
	assert.NotNil(t, client)

	assert.Equal(t, "", client.GetLoggerMsg())
}

func Test_TransToString(t *testing.T) {
	x := map[string]interface{}{
		"foo": make(chan int),
	}
	str := TransToString(x)
	assert.Equal(t, str, "")
}

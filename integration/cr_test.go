package integration

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cr"

	"github.com/stretchr/testify/assert"

	"fmt"
	"os"
	"testing"
)

var crTestKey = "crtestkey"

func Test_CR_CreateNamespace(t *testing.T) {
	client, err := cr.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateCreateNamespaceRequest()
	request.SetDomain("cr.cn-hangzhou.aliyuncs.com")
	request.SetContentType("JSON")
	content := fmt.Sprintf(
		`{
			"Namespace":{
				"Namespace":"%s"
			}
		}`, crTestKey,
	)
	request.SetContent([]byte(content))

	response, err := client.CreateNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_CR_UpdateNamespace(t *testing.T) {
	client, err := cr.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateUpdateNamespaceRequest()
	request.SetDomain("cr.cn-hangzhou.aliyuncs.com")
	request.Namespace = crTestKey
	request.SetContentType("JSON")
	content := fmt.Sprintf(
		`{
			"Namespace":{
				"AutoCreate":%v,
				"DefaultVisibility":"%s"
			}
		}`, false, "PUBLIC",
	)
	request.SetContent([]byte(content))

	response, err := client.UpdateNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_CR_GetNamespace(t *testing.T) {
	client, err := cr.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateGetNamespaceRequest()
	request.SetDomain("cr.cn-hangzhou.aliyuncs.com")
	request.Namespace = crTestKey

	response, err := client.GetNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_CR_GetNamespaceList(t *testing.T) {
	client, err := cr.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateGetNamespaceListRequest()
	request.SetDomain("cr.cn-hangzhou.aliyuncs.com")

	response, err := client.GetNamespaceList(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_CR_DeleteNamespace(t *testing.T) {
	client, err := cr.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateDeleteNamespaceRequest()
	request.SetDomain("cr.cn-hangzhou.aliyuncs.com")
	request.Namespace = crTestKey

	response, err := client.DeleteNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

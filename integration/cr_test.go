package integration

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cr"

	"github.com/stretchr/testify/assert"

	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

var genKey func() string

func init() {
	rand.Seed(time.Now().UnixNano())
	genKey = genKeyFunc()
}

func genKeyFunc() func() string {
	var key int
	return func() string {
		if key == 0 {
			key = rand.Intn(1e8)
		}
		return fmt.Sprintf("%08d", key)
	}
}

func Test_CR_CreateNamespace(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateCreateNamespaceRequest()
	domain := fmt.Sprintf("cr." + os.Getenv("REGION_ID") + ".aliyuncs.com")
	request.SetDomain(domain)
	request.SetContentType("JSON")
	content := fmt.Sprintf(
		`{
			"Namespace":{
				"Namespace":"%s"
			}
		}`, genKey(),
	)
	request.SetContent([]byte(content))

	response, err := client.CreateNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	t.Log(content)
}

func Test_CR_UpdateNamespace(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateUpdateNamespaceRequest()
	domain := fmt.Sprintf("cr." + os.Getenv("REGION_ID") + ".aliyuncs.com")
	request.SetDomain(domain)
	request.Namespace = genKey()
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
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateGetNamespaceRequest()
	domain := fmt.Sprintf("cr." + os.Getenv("REGION_ID") + ".aliyuncs.com")
	request.SetDomain(domain)
	request.Namespace = genKey()

	response, err := client.GetNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_CR_GetNamespaceList(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateGetNamespaceListRequest()
	domain := fmt.Sprintf("cr." + os.Getenv("REGION_ID") + ".aliyuncs.com")
	request.SetDomain(domain)

	response, err := client.GetNamespaceList(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_CR_DeleteNamespace(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateDeleteNamespaceRequest()
	domain := fmt.Sprintf("cr." + os.Getenv("REGION_ID") + ".aliyuncs.com")
	request.SetDomain(domain)
	request.Namespace = genKey()

	response, err := client.DeleteNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

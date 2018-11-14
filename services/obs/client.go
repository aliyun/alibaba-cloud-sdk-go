package obs



import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
)

// Client is the sdk client struct, each func corresponds to an OpenAPI
type Client struct {
	sdk.Client
}

// NewClient creates a sdk client with environment variables
func NewClient() (client *Client, err error) {
	client = &Client{}
	err = client.Init()
	return
}

// NewClientWithOptions creates a sdk client with regionId/sdkConfig/credential
// this is the common api to create a sdk client
func NewClientWithOptions(regionId string, config *sdk.Config, credential auth.Credential) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithOptions(regionId, config, credential)
	return
}

// NewClientWithAccessKey is a shortcut to create sdk client with accesskey
// usage: https://help.aliyun.com/document_detail/66217.html
func NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithAccessKey(regionId, accessKeyId, accessKeySecret)
	return
}

// NewClientWithStsToken is a shortcut to create sdk client with sts token
// usage: https://help.aliyun.com/document_detail/66222.html
func NewClientWithStsToken(regionId, stsAccessKeyId, stsAccessKeySecret, stsToken string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithStsToken(regionId, stsAccessKeyId, stsAccessKeySecret, stsToken)
	return
}

// NewClientWithRamRoleArn is a shortcut to create sdk client with ram roleArn
// usage: https://help.aliyun.com/document_detail/66222.html
func NewClientWithRamRoleArn(regionId string, accessKeyId, accessKeySecret, roleArn, roleSessionName string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithRamRoleArn(regionId, accessKeyId, accessKeySecret, roleArn, roleSessionName)
	return
}

// NewClientWithEcsRamRole is a shortcut to create sdk client with ecs ram role
// usage: https://help.aliyun.com/document_detail/66223.html
func NewClientWithEcsRamRole(regionId string, roleName string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithEcsRamRole(regionId, roleName)
	return
}

// NewClientWithRsaKeyPair is a shortcut to create sdk client with rsa key pair
// attention: rsa key pair auth is only Japan regions available
func NewClientWithRsaKeyPair(regionId string, publicKeyId, privateKey string, sessionExpiration int) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithRsaKeyPair(regionId, publicKeyId, privateKey, sessionExpiration)
	return
}

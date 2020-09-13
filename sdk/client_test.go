/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sdk

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials/provider"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/stretchr/testify/assert"
)

func Test_Client(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
		assert.Equal(t, "not support yet", err)
	}()
	NewClient()
}

func Test_NewClientWithOptions(t *testing.T) {
	c := NewConfig()
	credential := credentials.NewAccessKeyCredential("acesskeyid", "accesskeysecret")
	client, err := NewClientWithOptions("regionid", c, credential)
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func Test_NewClientWithAccessKey(t *testing.T) {
	client, err := NewClientWithAccessKey("regionid", "acesskeyid", "accesskeysecret")
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func Test_NewClientWithStsToken(t *testing.T) {
	client, err := NewClientWithStsToken("regionid", "acesskeyid", "accesskeysecret", "token")
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func Test_NewClientWithRamRoleArn(t *testing.T) {
	client, err := NewClientWithRamRoleArn("regionid", "acesskeyid", "accesskeysecret", "roleArn", "roleSessionName")
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func Test_NewClientWithEcsRamRole(t *testing.T) {
	client, err := NewClientWithEcsRamRole("regionid", "roleName")
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func Test_NewClientWithRsaKeyPair(t *testing.T) {
	client, err := NewClientWithRsaKeyPair("regionid", "publicKey", "privateKey", 3600)
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestInitWithProviderChain(t *testing.T) {

	// testcase1: No any environment variable
	c, err := NewClientWithProvider("cn-hangzhou")
	assert.NotNil(t, err)
	assert.Equal(t, &Client{}, c)
	assert.EqualError(t, err, "No credential found")

	// testcase2: AK
	os.Setenv(provider.ENVAccessKeyID, "AccessKeyId")
	os.Setenv(provider.ENVAccessKeySecret, "AccessKeySecret")

	c, err = NewClientWithProvider("cn-hangzhou")
	assert.Nil(t, err)
	expC, err := NewClientWithAccessKey("cn-hangzhou", "AccessKeyId", "AccessKeySecret")
	assert.Nil(t, err)
	c.httpDoHook = nil
	expC.httpDoHook = nil
	assert.Equal(t, expC, c)

	// testcase3:AK value is ""
	os.Setenv(provider.ENVAccessKeyID, "")
	os.Setenv(provider.ENVAccessKeySecret, "bbbb")
	c, err = NewClientWithProvider("cn-hangzhou")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "Environmental variable (ALIBABACLOUD_ACCESS_KEY_ID or ALIBABACLOUD_ACCESS_KEY_SECRET) is empty")
	assert.Equal(t, &Client{}, c)

	// testcase4: Profile value is ""
	os.Unsetenv(provider.ENVAccessKeyID)
	os.Unsetenv(provider.ENVAccessKeySecret)
	os.Setenv(provider.ENVCredentialFile, "")
	c, err = NewClientWithProvider("cn-hangzhou")
	assert.Equal(t, &Client{}, c)
	assert.EqualError(t, err, "Environment variable 'ALIBABA_CLOUD_CREDENTIALS_FILE' cannot be empty")

	// testcase5: Profile
	os.Setenv(provider.ENVCredentialFile, "./profile")
	c, err = NewClientWithProvider("cn-hangzhou")
	assert.Equal(t, &Client{}, c)
	assert.NotNil(t, err)
	// testcase6:Instances
	os.Unsetenv(provider.ENVCredentialFile)
	os.Setenv(provider.ENVEcsMetadata, "")
	c, err = NewClientWithProvider("cn-hangzhou")
	assert.Equal(t, &Client{}, c)
	assert.EqualError(t, err, "Environmental variable 'ALIBABA_CLOUD_ECS_METADATA' are empty")

	// testcase7: Custom Providers
	c, err = NewClientWithProvider("cn-hangzhou", provider.ProviderProfile, provider.ProviderEnv)
	assert.Equal(t, &Client{}, c)
	assert.EqualError(t, err, "No credential found")

}

func mockResponse(statusCode int, content string) (res *http.Response, err error) {
	status := strconv.Itoa(statusCode)
	res = &http.Response{
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		Header:     make(http.Header),
		StatusCode: statusCode,
		Status:     status + " " + http.StatusText(statusCode),
	}
	res.Body = ioutil.NopCloser(bytes.NewReader([]byte(content)))
	return
}

func Test_DoAction(t *testing.T) {
	client, err := NewClientWithAccessKey("regionid", "acesskeyid", "accesskeysecret")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, true, client.isRunning)
	request := requests.NewCommonRequest()
	request.Domain = "ecs.aliyuncs.com"
	request.Version = "2014-05-26"
	request.ApiName = "DescribeInstanceStatus"

	request.QueryParams["PageNumber"] = "1"
	request.QueryParams["PageSize"] = "30"
	request.TransToAcsRequest()
	response := responses.NewCommonResponse()
	client.SetHTTPDoHook(func(fn func(req *http.Request) (*http.Response, error)) func(req *http.Request) (*http.Response, error) {
		return func(req *http.Request) (*http.Response, error) {
			return mockResponse(200, "")
		}
	})
	err = client.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.GetHttpStatus())
	assert.Equal(t, "", response.GetHttpContentString())
	client.Shutdown()
	assert.Equal(t, false, client.isRunning)
}

func Test_ProcessCommonRequest(t *testing.T) {
	client, err := NewClientWithAccessKey("regionid", "acesskeyid", "accesskeysecret")
	assert.Nil(t, err)
	assert.NotNil(t, client)

	request := requests.NewCommonRequest()
	request.Domain = "ecs.aliyuncs.com"
	request.Version = "2014-05-26"
	request.ApiName = "DescribeInstanceStatus"

	request.QueryParams["PageNumber"] = "1"
	request.QueryParams["PageSize"] = "30"

	client.SetHTTPDoHook(func(fn func(req *http.Request) (*http.Response, error)) func(req *http.Request) (*http.Response, error) {
		return func(req *http.Request) (*http.Response, error) {
			return mockResponse(200, "")
		}
	})
	response, err := client.ProcessCommonRequest(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.GetHttpStatus())
	assert.Equal(t, "", response.GetHttpContentString())
}

func Test_DoAction_With500(t *testing.T) {
	client, err := NewClientWithAccessKey("regionid", "acesskeyid", "accesskeysecret")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, true, client.isRunning)
	request := requests.NewCommonRequest()
	request.Domain = "ecs.aliyuncs.com"
	request.Version = "2014-05-26"
	request.ApiName = "DescribeInstanceStatus"

	request.QueryParams["PageNumber"] = "1"
	request.QueryParams["PageSize"] = "30"
	request.TransToAcsRequest()
	response := responses.NewCommonResponse()
	client.SetHTTPDoHook(func(fn func(req *http.Request) (*http.Response, error)) func(req *http.Request) (*http.Response, error) {
		return func(req *http.Request) (*http.Response, error) {
			return mockResponse(500, "Server Internel Error")
		}
	})
	err = client.DoAction(request, response)
	assert.NotNil(t, err)
	assert.Equal(t, "SDK.ServerError\nErrorCode: \nRecommend: \nRequestId: \nMessage: Server Internel Error", err.Error())
	assert.Equal(t, 500, response.GetHttpStatus())
	assert.Equal(t, "Server Internel Error", response.GetHttpContentString())
}

// func Test_EnableAsync(t *testing.T) {
// 	client, err := NewClientWithAccessKey("regionid", "acesskeyid", "accesskeysecret")
// 	assert.Nil(t, err)
// 	assert.NotNil(t, client)
// 	assert.Equal(t, true, client.isRunning)
// 	client.EnableAsync(2, 8)
// 	client.Shutdown()
// 	assert.Equal(t, false, client.isRunning)
// }

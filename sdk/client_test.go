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
	"crypto/tls"
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"testing"
)

var client, clientKeyPair, clientEcs, clientRoleArn *Client

type TestConfig struct {
	AccessKeyId     string
	AccessKeySecret string
	PublicKeyId     string
	PrivateKey      string
	RoleArn         string
	ChildAK         string
	ChildSecret     string
}

type MockResponse struct {
	Headers     map[string]string
	Body        string
	Params      map[string]string
	RemoteAddr  string
	RemoteHost  string
	QueryString string
	RequestURL  string
}

func TestMain(m *testing.M) {
	testSetup()
	result := m.Run()
	testTearDown()
	os.Exit(result)
}

func testSetup() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	configFilePath := usr.HomeDir + "/aliyun-sdk.json"
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	var config TestConfig
	json.Unmarshal(data, &config)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &Client{
		config: &Config{
			HttpTransport: tr,
		},
	}
	clientEcs = &Client{}
	clientRoleArn = &Client{
		config: &Config{
			HttpTransport: tr,
		},
	}
	clientKeyPair = &Client{
		config: &Config{
			HttpTransport: tr,
		},
	}
	err = client.InitWithAccessKey("cn-hangzhou", config.AccessKeyId, config.AccessKeySecret)
	err = clientKeyPair.InitWithKeyPair("cn-hangzhou", config.PublicKeyId, config.PrivateKey, 3600)
	err = clientEcs.InitWithEcsInstance("cn-hangzhou", "conan")
	err = clientRoleArn.InitWithRoleArn("cn-hangzhou", config.ChildAK, config.ChildSecret, config.RoleArn, "clientTest")
}

func testTearDown() {

}

func TestNewClientWithAccessKey(t *testing.T) {
	assert.NotNil(t, client, "NewClientWithAccessKey failed")
}

func TestRoaGet(t *testing.T) {
	request := getFtTestRoaRequest()

	response := &responses.BaseResponse{}
	err := client.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
	assert.Equal(t, "HeaderParamValue", responseBean.Headers["Header-Param"])
}

func TestRoaPostForm(t *testing.T) {
	request := getFtTestRoaRequest()
	request.Method = requests.POST
	request.FormParams["BodyParam"] = "BodyParamValue"

	response := &responses.BaseResponse{}
	err := client.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
	assert.Equal(t, "HeaderParamValue", responseBean.Headers["Header-Param"])
	assert.Equal(t, "BodyParamValue", responseBean.Params["BodyParam"])
}

func TestRoaPostStream(t *testing.T) {
	request := getFtTestRoaRequest()
	request.Method = requests.POST
	request.Content = []byte("TestContent")

	response := &responses.BaseResponse{}
	err := client.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
	assert.Equal(t, "HeaderParamValue", responseBean.Headers["Header-Param"])
	assert.Equal(t, "TestContent", responseBean.Body)
}

func TestRoaPostJson(t *testing.T) {
	request := getFtTestRoaRequest()
	request.Method = requests.POST
	dataMap := map[string]string{"key": "value"}
	data, err := json.Marshal(dataMap)
	assert.Nil(t, err)
	request.Content = data
	request.SetContentType(requests.Json)

	response := &responses.BaseResponse{}
	err = client.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
	assert.Equal(t, "HeaderParamValue", responseBean.Headers["Header-Param"])
	assert.Equal(t, requests.Json, responseBean.Headers["Content-Type"])
	assert.Equal(t, string(data), responseBean.Body)
}

func TestRpcGet(t *testing.T) {
	request := getFtTestRpcRequest()
	request.Method = requests.GET

	response := &responses.BaseResponse{}
	err := client.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
}

func TestRpcGetForHttps(t *testing.T) {
	request := getFtTestRpcRequest()
	request.Method = requests.GET
	request.Scheme = requests.HTTPS

	response := &responses.BaseResponse{}
	err := client.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
}

func TestRoaGetForHttps(t *testing.T) {
	request := getFtTestRoaRequest()
	request.Scheme = requests.HTTPS

	response := &responses.BaseResponse{}
	err := client.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
	assert.Equal(t, "HeaderParamValue", responseBean.Headers["Header-Param"])
}

func TestRpcPost(t *testing.T) {
	request := getFtTestRpcRequest()
	request.FormParams["BodyParam"] = "BodyParamValue"

	response := &responses.BaseResponse{}
	err := client.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
	assert.Equal(t, "BodyParamValue", responseBean.Params["BodyParam"])
}

func getFtTestRoaRequest() (request *requests.RoaRequest) {
	request = &requests.RoaRequest{}
	request.InitWithApiInfo("Ft", "2016-01-02", "TestRoaApi", "/web/cloudapi", "", "")
	request.Domain = "ft.aliyuncs.com"

	request.Headers["Header-Param"] = "HeaderParamValue"
	request.QueryParams["QueryParam"] = "QueryParamValue"

	return
}

func getFtTestRpcRequest() (request *requests.RpcRequest) {
	request = &requests.RpcRequest{}
	request.InitWithApiInfo("Ft", "2016-01-01", "TestRpcApi", "", "")
	request.Domain = "ft.aliyuncs.com"
	request.QueryParams["QueryParam"] = "QueryParamValue"
	return
}

func getFtTestRpcRequestForEndpointLocation() (request *requests.RpcRequest) {
	request = &requests.RpcRequest{}
	request.InitWithApiInfo("Ft", "2016-01-01", "TestRpcApi", "ft", "openAPI")
	request.RegionId = "cn-hangzhou"
	request.QueryParams["QueryParam"] = "QueryParamValue"
	return
}

func TestCommonRpcRequest(t *testing.T) {
	rpcRequest := requests.NewCommonRequest()
	rpcRequest.Product = "Ft"
	rpcRequest.Version = "2016-01-01"
	rpcRequest.Domain = "ft.aliyuncs.com"
	rpcRequest.ApiName = "TestRpcApi"
	rpcRequest.Method = "POST"

	rpcRequest.QueryParams["QueryParam"] = "QueryParamValue"
	rpcRequest.FormParams["BodyParam"] = "BodyParamValue"

	response, err := client.ProcessCommonRequest(rpcRequest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
	assert.Equal(t, "BodyParamValue", responseBean.Params["BodyParam"])
}

func TestCommonRoaRequest(t *testing.T) {
	roaRequest := requests.NewCommonRequest()
	roaRequest.Product = "Ft"
	roaRequest.Version = "2016-01-02"
	roaRequest.ApiName = "TestRoaApi"
	roaRequest.PathPattern = "/web/cloudapi"
	roaRequest.Domain = "ft.aliyuncs.com"
	roaRequest.Method = "POST"

	roaRequest.QueryParams["QueryParam"] = "QueryParamValue"
	roaRequest.FormParams["BodyParam"] = "BodyParamValue"

	response, err := client.ProcessCommonRequest(roaRequest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
	assert.Equal(t, "BodyParamValue", responseBean.Params["BodyParam"])
}

func TestRpcGetForEndpointXml(t *testing.T) {
	request := getFtTestRpcRequestForEndpointLocation()
	request.Method = requests.GET
	request.RegionId = "cn-shanghai"

	response := &responses.BaseResponse{}
	err := client.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
}

func TestRpcGetForLocation(t *testing.T) {
	request := getFtTestRpcRequestForEndpointLocation()
	request.Method = requests.GET

	response := &responses.BaseResponse{}
	err := client.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
}

func TestRpcGetForLocationCache(t *testing.T) {
	request := getFtTestRpcRequestForEndpointLocation()
	request.Method = requests.GET

	response := &responses.BaseResponse{}
	err := client.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])

	request2 := getFtTestRpcRequestForEndpointLocation()
	request2.Method = requests.GET
	err = client.DoAction(request2, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
}

func TestRpcGetForKeyPair(t *testing.T) {
	request := getFtTestRpcRequest()
	request.Method = requests.GET

	response := &responses.BaseResponse{}
	err := clientKeyPair.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
}

/*func TestRpcGetForEcs(t *testing.T) {
	//测试接口，想测试的时候，要替换掉singer_ecs_instance中对应的变量，并且还要提供一个mock服务
	//requestUrl := "http://localhost:3500/latest/meta-data/ram/security-credentials/roleNameTest.json"
	request := getFtTestRpcRequest()
	request.Method = requests.GET

	response := &responses.BaseResponse{}
	err := clientEcs.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])

	err = clientEcs.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
}*/

func TestRpcGetForRoleArn(t *testing.T) {
	request := getFtTestRpcRequest()
	request.Method = requests.GET

	response := &responses.BaseResponse{}
	err := clientRoleArn.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	var responseBean MockResponse
	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])

	err = clientRoleArn.DoAction(request, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.GetHttpStatus(), response.GetHttpContentString())
	assert.NotNil(t, response.GetHttpContentString())

	json.Unmarshal([]byte(response.GetHttpContentString()), &responseBean)

	assert.Equal(t, "QueryParamValue", responseBean.Params["QueryParam"])
}

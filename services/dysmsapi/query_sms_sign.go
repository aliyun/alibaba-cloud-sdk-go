package dysmsapi

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// QuerySmsSign invokes the dysmsapi.QuerySmsSign API synchronously
func (client *Client) QuerySmsSign(request *QuerySmsSignRequest) (response *QuerySmsSignResponse, err error) {
	response = CreateQuerySmsSignResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySmsSignWithChan invokes the dysmsapi.QuerySmsSign API asynchronously
func (client *Client) QuerySmsSignWithChan(request *QuerySmsSignRequest) (<-chan *QuerySmsSignResponse, <-chan error) {
	responseChan := make(chan *QuerySmsSignResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySmsSign(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// QuerySmsSignWithCallback invokes the dysmsapi.QuerySmsSign API asynchronously
func (client *Client) QuerySmsSignWithCallback(request *QuerySmsSignRequest, callback func(response *QuerySmsSignResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySmsSignResponse
		var err error
		defer close(result)
		response, err = client.QuerySmsSign(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// QuerySmsSignRequest is the request struct for api QuerySmsSign
type QuerySmsSignRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SignName             string           `position:"Query" name:"SignName"`
}

// QuerySmsSignResponse is the response struct for api QuerySmsSign
type QuerySmsSignResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	SignStatus int    `json:"SignStatus" xml:"SignStatus"`
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	CreateDate string `json:"CreateDate" xml:"CreateDate"`
	Reason     string `json:"Reason" xml:"Reason"`
	SignName   string `json:"SignName" xml:"SignName"`
}

// CreateQuerySmsSignRequest creates a request to invoke QuerySmsSign API
func CreateQuerySmsSignRequest() (request *QuerySmsSignRequest) {
	request = &QuerySmsSignRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "QuerySmsSign", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySmsSignResponse creates a response to parse from QuerySmsSign response
func CreateQuerySmsSignResponse() (response *QuerySmsSignResponse) {
	response = &QuerySmsSignResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

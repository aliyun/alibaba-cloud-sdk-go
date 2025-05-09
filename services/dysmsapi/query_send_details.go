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

// QuerySendDetails invokes the dysmsapi.QuerySendDetails API synchronously
func (client *Client) QuerySendDetails(request *QuerySendDetailsRequest) (response *QuerySendDetailsResponse, err error) {
	response = CreateQuerySendDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySendDetailsWithChan invokes the dysmsapi.QuerySendDetails API asynchronously
func (client *Client) QuerySendDetailsWithChan(request *QuerySendDetailsRequest) (<-chan *QuerySendDetailsResponse, <-chan error) {
	responseChan := make(chan *QuerySendDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySendDetails(request)
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

// QuerySendDetailsWithCallback invokes the dysmsapi.QuerySendDetails API asynchronously
func (client *Client) QuerySendDetailsWithCallback(request *QuerySendDetailsRequest, callback func(response *QuerySendDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySendDetailsResponse
		var err error
		defer close(result)
		response, err = client.QuerySendDetails(request)
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

// QuerySendDetailsRequest is the request struct for api QuerySendDetails
type QuerySendDetailsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PhoneNumber          string           `position:"Query" name:"PhoneNumber"`
	SendDate             string           `position:"Query" name:"SendDate"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	CurrentPage          requests.Integer `position:"Query" name:"CurrentPage"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	BizId                string           `position:"Query" name:"BizId"`
}

// QuerySendDetailsResponse is the response struct for api QuerySendDetails
type QuerySendDetailsResponse struct {
	*responses.BaseResponse
	Code              string            `json:"Code" xml:"Code"`
	Message           string            `json:"Message" xml:"Message"`
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	TotalCount        string            `json:"TotalCount" xml:"TotalCount"`
	SmsSendDetailDTOs SmsSendDetailDTOs `json:"SmsSendDetailDTOs" xml:"SmsSendDetailDTOs"`
}

// CreateQuerySendDetailsRequest creates a request to invoke QuerySendDetails API
func CreateQuerySendDetailsRequest() (request *QuerySendDetailsRequest) {
	request = &QuerySendDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "QuerySendDetails", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySendDetailsResponse creates a response to parse from QuerySendDetails response
func CreateQuerySendDetailsResponse() (response *QuerySendDetailsResponse) {
	response = &QuerySendDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

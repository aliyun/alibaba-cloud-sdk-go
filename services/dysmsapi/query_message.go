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

// QueryMessage invokes the dysmsapi.QueryMessage API synchronously
func (client *Client) QueryMessage(request *QueryMessageRequest) (response *QueryMessageResponse, err error) {
	response = CreateQueryMessageResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMessageWithChan invokes the dysmsapi.QueryMessage API asynchronously
func (client *Client) QueryMessageWithChan(request *QueryMessageRequest) (<-chan *QueryMessageResponse, <-chan error) {
	responseChan := make(chan *QueryMessageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMessage(request)
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

// QueryMessageWithCallback invokes the dysmsapi.QueryMessage API asynchronously
func (client *Client) QueryMessageWithCallback(request *QueryMessageRequest, callback func(response *QueryMessageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMessageResponse
		var err error
		defer close(result)
		response, err = client.QueryMessage(request)
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

// QueryMessageRequest is the request struct for api QueryMessage
type QueryMessageRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	MessageId            string           `position:"Query" name:"MessageId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryMessageResponse is the response struct for api QueryMessage
type QueryMessageResponse struct {
	*responses.BaseResponse
	Status              string       `json:"Status" xml:"Status"`
	ErrorDescription    string       `json:"ErrorDescription" xml:"ErrorDescription"`
	ResponseCode        string       `json:"ResponseCode" xml:"ResponseCode"`
	ReceiveDate         string       `json:"ReceiveDate" xml:"ReceiveDate"`
	Message             string       `json:"Message" xml:"Message"`
	ResponseDescription string       `json:"ResponseDescription" xml:"ResponseDescription"`
	ErrorCode           string       `json:"ErrorCode" xml:"ErrorCode"`
	SendDate            string       `json:"SendDate" xml:"SendDate"`
	To                  string       `json:"To" xml:"To"`
	MessageId           string       `json:"MessageId" xml:"MessageId"`
	RequestId           string       `json:"RequestId" xml:"RequestId"`
	NumberDetail        NumberDetail `json:"NumberDetail" xml:"NumberDetail"`
}

// CreateQueryMessageRequest creates a request to invoke QueryMessage API
func CreateQueryMessageRequest() (request *QueryMessageRequest) {
	request = &QueryMessageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2018-05-01", "QueryMessage", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryMessageResponse creates a response to parse from QueryMessage response
func CreateQueryMessageResponse() (response *QueryMessageResponse) {
	response = &QueryMessageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

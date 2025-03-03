package resourcemanager

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

// SendVerificationCodeForBindSecureMobilePhone invokes the resourcemanager.SendVerificationCodeForBindSecureMobilePhone API synchronously
func (client *Client) SendVerificationCodeForBindSecureMobilePhone(request *SendVerificationCodeForBindSecureMobilePhoneRequest) (response *SendVerificationCodeForBindSecureMobilePhoneResponse, err error) {
	response = CreateSendVerificationCodeForBindSecureMobilePhoneResponse()
	err = client.DoAction(request, response)
	return
}

// SendVerificationCodeForBindSecureMobilePhoneWithChan invokes the resourcemanager.SendVerificationCodeForBindSecureMobilePhone API asynchronously
func (client *Client) SendVerificationCodeForBindSecureMobilePhoneWithChan(request *SendVerificationCodeForBindSecureMobilePhoneRequest) (<-chan *SendVerificationCodeForBindSecureMobilePhoneResponse, <-chan error) {
	responseChan := make(chan *SendVerificationCodeForBindSecureMobilePhoneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendVerificationCodeForBindSecureMobilePhone(request)
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

// SendVerificationCodeForBindSecureMobilePhoneWithCallback invokes the resourcemanager.SendVerificationCodeForBindSecureMobilePhone API asynchronously
func (client *Client) SendVerificationCodeForBindSecureMobilePhoneWithCallback(request *SendVerificationCodeForBindSecureMobilePhoneRequest, callback func(response *SendVerificationCodeForBindSecureMobilePhoneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendVerificationCodeForBindSecureMobilePhoneResponse
		var err error
		defer close(result)
		response, err = client.SendVerificationCodeForBindSecureMobilePhone(request)
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

// SendVerificationCodeForBindSecureMobilePhoneRequest is the request struct for api SendVerificationCodeForBindSecureMobilePhone
type SendVerificationCodeForBindSecureMobilePhoneRequest struct {
	*requests.RpcRequest
	SecureMobilePhone string `position:"Query" name:"SecureMobilePhone"`
	AccountId         string `position:"Query" name:"AccountId"`
}

// SendVerificationCodeForBindSecureMobilePhoneResponse is the response struct for api SendVerificationCodeForBindSecureMobilePhone
type SendVerificationCodeForBindSecureMobilePhoneResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ExpirationDate string `json:"ExpirationDate" xml:"ExpirationDate"`
}

// CreateSendVerificationCodeForBindSecureMobilePhoneRequest creates a request to invoke SendVerificationCodeForBindSecureMobilePhone API
func CreateSendVerificationCodeForBindSecureMobilePhoneRequest() (request *SendVerificationCodeForBindSecureMobilePhoneRequest) {
	request = &SendVerificationCodeForBindSecureMobilePhoneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "SendVerificationCodeForBindSecureMobilePhone", "resourcemanager", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSendVerificationCodeForBindSecureMobilePhoneResponse creates a response to parse from SendVerificationCodeForBindSecureMobilePhone response
func CreateSendVerificationCodeForBindSecureMobilePhoneResponse() (response *SendVerificationCodeForBindSecureMobilePhoneResponse) {
	response = &SendVerificationCodeForBindSecureMobilePhoneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

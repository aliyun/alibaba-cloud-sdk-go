package arms

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

// DeleteNotificationPolicy invokes the arms.DeleteNotificationPolicy API synchronously
func (client *Client) DeleteNotificationPolicy(request *DeleteNotificationPolicyRequest) (response *DeleteNotificationPolicyResponse, err error) {
	response = CreateDeleteNotificationPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteNotificationPolicyWithChan invokes the arms.DeleteNotificationPolicy API asynchronously
func (client *Client) DeleteNotificationPolicyWithChan(request *DeleteNotificationPolicyRequest) (<-chan *DeleteNotificationPolicyResponse, <-chan error) {
	responseChan := make(chan *DeleteNotificationPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteNotificationPolicy(request)
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

// DeleteNotificationPolicyWithCallback invokes the arms.DeleteNotificationPolicy API asynchronously
func (client *Client) DeleteNotificationPolicyWithCallback(request *DeleteNotificationPolicyRequest, callback func(response *DeleteNotificationPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteNotificationPolicyResponse
		var err error
		defer close(result)
		response, err = client.DeleteNotificationPolicy(request)
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

// DeleteNotificationPolicyRequest is the request struct for api DeleteNotificationPolicy
type DeleteNotificationPolicyRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Query" name:"Id"`
}

// DeleteNotificationPolicyResponse is the response struct for api DeleteNotificationPolicy
type DeleteNotificationPolicyResponse struct {
	*responses.BaseResponse
	RequestId                         string `json:"RequestId" xml:"RequestId"`
	DeleteNotificationPolicyIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
}

// CreateDeleteNotificationPolicyRequest creates a request to invoke DeleteNotificationPolicy API
func CreateDeleteNotificationPolicyRequest() (request *DeleteNotificationPolicyRequest) {
	request = &DeleteNotificationPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "DeleteNotificationPolicy", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteNotificationPolicyResponse creates a response to parse from DeleteNotificationPolicy response
func CreateDeleteNotificationPolicyResponse() (response *DeleteNotificationPolicyResponse) {
	response = &DeleteNotificationPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

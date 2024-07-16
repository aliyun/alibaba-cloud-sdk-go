package swas_open

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

// RemoveCustomImageShareAccount invokes the swas_open.RemoveCustomImageShareAccount API synchronously
func (client *Client) RemoveCustomImageShareAccount(request *RemoveCustomImageShareAccountRequest) (response *RemoveCustomImageShareAccountResponse, err error) {
	response = CreateRemoveCustomImageShareAccountResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveCustomImageShareAccountWithChan invokes the swas_open.RemoveCustomImageShareAccount API asynchronously
func (client *Client) RemoveCustomImageShareAccountWithChan(request *RemoveCustomImageShareAccountRequest) (<-chan *RemoveCustomImageShareAccountResponse, <-chan error) {
	responseChan := make(chan *RemoveCustomImageShareAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveCustomImageShareAccount(request)
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

// RemoveCustomImageShareAccountWithCallback invokes the swas_open.RemoveCustomImageShareAccount API asynchronously
func (client *Client) RemoveCustomImageShareAccountWithCallback(request *RemoveCustomImageShareAccountRequest, callback func(response *RemoveCustomImageShareAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveCustomImageShareAccountResponse
		var err error
		defer close(result)
		response, err = client.RemoveCustomImageShareAccount(request)
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

// RemoveCustomImageShareAccountRequest is the request struct for api RemoveCustomImageShareAccount
type RemoveCustomImageShareAccountRequest struct {
	*requests.RpcRequest
	ImageId     string    `position:"Query" name:"ImageId"`
	ClientToken string    `position:"Query" name:"ClientToken"`
	Account     *[]string `position:"Query" name:"Account"  type:"Repeated"`
}

// RemoveCustomImageShareAccountResponse is the response struct for api RemoveCustomImageShareAccount
type RemoveCustomImageShareAccountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveCustomImageShareAccountRequest creates a request to invoke RemoveCustomImageShareAccount API
func CreateRemoveCustomImageShareAccountRequest() (request *RemoveCustomImageShareAccountRequest) {
	request = &RemoveCustomImageShareAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "RemoveCustomImageShareAccount", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveCustomImageShareAccountResponse creates a response to parse from RemoveCustomImageShareAccount response
func CreateRemoveCustomImageShareAccountResponse() (response *RemoveCustomImageShareAccountResponse) {
	response = &RemoveCustomImageShareAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

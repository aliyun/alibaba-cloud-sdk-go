package cloud_siem

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

// BindAccount invokes the cloud_siem.BindAccount API synchronously
func (client *Client) BindAccount(request *BindAccountRequest) (response *BindAccountResponse, err error) {
	response = CreateBindAccountResponse()
	err = client.DoAction(request, response)
	return
}

// BindAccountWithChan invokes the cloud_siem.BindAccount API asynchronously
func (client *Client) BindAccountWithChan(request *BindAccountRequest) (<-chan *BindAccountResponse, <-chan error) {
	responseChan := make(chan *BindAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindAccount(request)
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

// BindAccountWithCallback invokes the cloud_siem.BindAccount API asynchronously
func (client *Client) BindAccountWithCallback(request *BindAccountRequest, callback func(response *BindAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindAccountResponse
		var err error
		defer close(result)
		response, err = client.BindAccount(request)
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

// BindAccountRequest is the request struct for api BindAccount
type BindAccountRequest struct {
	*requests.RpcRequest
	CloudCode   string `position:"Body" name:"CloudCode"`
	AccountId   string `position:"Body" name:"AccountId"`
	AccountName string `position:"Body" name:"AccountName"`
	AccessId    string `position:"Body" name:"AccessId"`
}

// BindAccountResponse is the response struct for api BindAccount
type BindAccountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateBindAccountRequest creates a request to invoke BindAccount API
func CreateBindAccountRequest() (request *BindAccountRequest) {
	request = &BindAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "BindAccount", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBindAccountResponse creates a response to parse from BindAccount response
func CreateBindAccountResponse() (response *BindAccountResponse) {
	response = &BindAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

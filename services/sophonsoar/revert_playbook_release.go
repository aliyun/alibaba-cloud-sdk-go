package sophonsoar

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

// RevertPlaybookRelease invokes the sophonsoar.RevertPlaybookRelease API synchronously
func (client *Client) RevertPlaybookRelease(request *RevertPlaybookReleaseRequest) (response *RevertPlaybookReleaseResponse, err error) {
	response = CreateRevertPlaybookReleaseResponse()
	err = client.DoAction(request, response)
	return
}

// RevertPlaybookReleaseWithChan invokes the sophonsoar.RevertPlaybookRelease API asynchronously
func (client *Client) RevertPlaybookReleaseWithChan(request *RevertPlaybookReleaseRequest) (<-chan *RevertPlaybookReleaseResponse, <-chan error) {
	responseChan := make(chan *RevertPlaybookReleaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevertPlaybookRelease(request)
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

// RevertPlaybookReleaseWithCallback invokes the sophonsoar.RevertPlaybookRelease API asynchronously
func (client *Client) RevertPlaybookReleaseWithCallback(request *RevertPlaybookReleaseRequest, callback func(response *RevertPlaybookReleaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevertPlaybookReleaseResponse
		var err error
		defer close(result)
		response, err = client.RevertPlaybookRelease(request)
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

// RevertPlaybookReleaseRequest is the request struct for api RevertPlaybookRelease
type RevertPlaybookReleaseRequest struct {
	*requests.RpcRequest
	RoleFor       string           `position:"Query" name:"RoleFor"`
	PlaybookUuid  string           `position:"Body" name:"PlaybookUuid"`
	PlayReleaseId requests.Integer `position:"Body" name:"PlayReleaseId"`
	IsPublish     requests.Boolean `position:"Body" name:"IsPublish"`
	RoleType      string           `position:"Query" name:"RoleType"`
	Lang          string           `position:"Body" name:"Lang"`
}

// RevertPlaybookReleaseResponse is the response struct for api RevertPlaybookRelease
type RevertPlaybookReleaseResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRevertPlaybookReleaseRequest creates a request to invoke RevertPlaybookRelease API
func CreateRevertPlaybookReleaseRequest() (request *RevertPlaybookReleaseRequest) {
	request = &RevertPlaybookReleaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sophonsoar", "2022-07-28", "RevertPlaybookRelease", "", "")
	request.Method = requests.POST
	return
}

// CreateRevertPlaybookReleaseResponse creates a response to parse from RevertPlaybookRelease response
func CreateRevertPlaybookReleaseResponse() (response *RevertPlaybookReleaseResponse) {
	response = &RevertPlaybookReleaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package csas

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

// ListApplicationsForPrivateAccessTag invokes the csas.ListApplicationsForPrivateAccessTag API synchronously
func (client *Client) ListApplicationsForPrivateAccessTag(request *ListApplicationsForPrivateAccessTagRequest) (response *ListApplicationsForPrivateAccessTagResponse, err error) {
	response = CreateListApplicationsForPrivateAccessTagResponse()
	err = client.DoAction(request, response)
	return
}

// ListApplicationsForPrivateAccessTagWithChan invokes the csas.ListApplicationsForPrivateAccessTag API asynchronously
func (client *Client) ListApplicationsForPrivateAccessTagWithChan(request *ListApplicationsForPrivateAccessTagRequest) (<-chan *ListApplicationsForPrivateAccessTagResponse, <-chan error) {
	responseChan := make(chan *ListApplicationsForPrivateAccessTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListApplicationsForPrivateAccessTag(request)
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

// ListApplicationsForPrivateAccessTagWithCallback invokes the csas.ListApplicationsForPrivateAccessTag API asynchronously
func (client *Client) ListApplicationsForPrivateAccessTagWithCallback(request *ListApplicationsForPrivateAccessTagRequest, callback func(response *ListApplicationsForPrivateAccessTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListApplicationsForPrivateAccessTagResponse
		var err error
		defer close(result)
		response, err = client.ListApplicationsForPrivateAccessTag(request)
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

// ListApplicationsForPrivateAccessTagRequest is the request struct for api ListApplicationsForPrivateAccessTag
type ListApplicationsForPrivateAccessTagRequest struct {
	*requests.RpcRequest
	TagIds   *[]string `position:"Query" name:"TagIds"  type:"Repeated"`
	SourceIp string    `position:"Query" name:"SourceIp"`
}

// ListApplicationsForPrivateAccessTagResponse is the response struct for api ListApplicationsForPrivateAccessTag
type ListApplicationsForPrivateAccessTagResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Tags      []Tag  `json:"Tags" xml:"Tags"`
}

// CreateListApplicationsForPrivateAccessTagRequest creates a request to invoke ListApplicationsForPrivateAccessTag API
func CreateListApplicationsForPrivateAccessTagRequest() (request *ListApplicationsForPrivateAccessTagRequest) {
	request = &ListApplicationsForPrivateAccessTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "ListApplicationsForPrivateAccessTag", "", "")
	request.Method = requests.GET
	return
}

// CreateListApplicationsForPrivateAccessTagResponse creates a response to parse from ListApplicationsForPrivateAccessTag response
func CreateListApplicationsForPrivateAccessTagResponse() (response *ListApplicationsForPrivateAccessTagResponse) {
	response = &ListApplicationsForPrivateAccessTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package live

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

// DeleteLiveStreamMerge invokes the live.DeleteLiveStreamMerge API synchronously
func (client *Client) DeleteLiveStreamMerge(request *DeleteLiveStreamMergeRequest) (response *DeleteLiveStreamMergeResponse, err error) {
	response = CreateDeleteLiveStreamMergeResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveStreamMergeWithChan invokes the live.DeleteLiveStreamMerge API asynchronously
func (client *Client) DeleteLiveStreamMergeWithChan(request *DeleteLiveStreamMergeRequest) (<-chan *DeleteLiveStreamMergeResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveStreamMergeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveStreamMerge(request)
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

// DeleteLiveStreamMergeWithCallback invokes the live.DeleteLiveStreamMerge API asynchronously
func (client *Client) DeleteLiveStreamMergeWithCallback(request *DeleteLiveStreamMergeRequest, callback func(response *DeleteLiveStreamMergeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveStreamMergeResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveStreamMerge(request)
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

// DeleteLiveStreamMergeRequest is the request struct for api DeleteLiveStreamMerge
type DeleteLiveStreamMergeRequest struct {
	*requests.RpcRequest
	AppName    string           `position:"Query" name:"AppName"`
	StreamName string           `position:"Query" name:"StreamName"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteLiveStreamMergeResponse is the response struct for api DeleteLiveStreamMerge
type DeleteLiveStreamMergeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLiveStreamMergeRequest creates a request to invoke DeleteLiveStreamMerge API
func CreateDeleteLiveStreamMergeRequest() (request *DeleteLiveStreamMergeRequest) {
	request = &DeleteLiveStreamMergeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteLiveStreamMerge", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLiveStreamMergeResponse creates a response to parse from DeleteLiveStreamMerge response
func CreateDeleteLiveStreamMergeResponse() (response *DeleteLiveStreamMergeResponse) {
	response = &DeleteLiveStreamMergeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

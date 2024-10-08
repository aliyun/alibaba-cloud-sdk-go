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

// UnbanLiveMessageGroup invokes the live.UnbanLiveMessageGroup API synchronously
func (client *Client) UnbanLiveMessageGroup(request *UnbanLiveMessageGroupRequest) (response *UnbanLiveMessageGroupResponse, err error) {
	response = CreateUnbanLiveMessageGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UnbanLiveMessageGroupWithChan invokes the live.UnbanLiveMessageGroup API asynchronously
func (client *Client) UnbanLiveMessageGroupWithChan(request *UnbanLiveMessageGroupRequest) (<-chan *UnbanLiveMessageGroupResponse, <-chan error) {
	responseChan := make(chan *UnbanLiveMessageGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbanLiveMessageGroup(request)
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

// UnbanLiveMessageGroupWithCallback invokes the live.UnbanLiveMessageGroup API asynchronously
func (client *Client) UnbanLiveMessageGroupWithCallback(request *UnbanLiveMessageGroupRequest, callback func(response *UnbanLiveMessageGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbanLiveMessageGroupResponse
		var err error
		defer close(result)
		response, err = client.UnbanLiveMessageGroup(request)
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

// UnbanLiveMessageGroupRequest is the request struct for api UnbanLiveMessageGroup
type UnbanLiveMessageGroupRequest struct {
	*requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	DataCenter string `position:"Query" name:"DataCenter"`
	AppId      string `position:"Query" name:"AppId"`
}

// UnbanLiveMessageGroupResponse is the response struct for api UnbanLiveMessageGroup
type UnbanLiveMessageGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnbanLiveMessageGroupRequest creates a request to invoke UnbanLiveMessageGroup API
func CreateUnbanLiveMessageGroupRequest() (request *UnbanLiveMessageGroupRequest) {
	request = &UnbanLiveMessageGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UnbanLiveMessageGroup", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnbanLiveMessageGroupResponse creates a response to parse from UnbanLiveMessageGroup response
func CreateUnbanLiveMessageGroupResponse() (response *UnbanLiveMessageGroupResponse) {
	response = &UnbanLiveMessageGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

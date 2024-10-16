package gwlb

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

// RemoveServersFromServerGroup invokes the gwlb.RemoveServersFromServerGroup API synchronously
func (client *Client) RemoveServersFromServerGroup(request *RemoveServersFromServerGroupRequest) (response *RemoveServersFromServerGroupResponse, err error) {
	response = CreateRemoveServersFromServerGroupResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveServersFromServerGroupWithChan invokes the gwlb.RemoveServersFromServerGroup API asynchronously
func (client *Client) RemoveServersFromServerGroupWithChan(request *RemoveServersFromServerGroupRequest) (<-chan *RemoveServersFromServerGroupResponse, <-chan error) {
	responseChan := make(chan *RemoveServersFromServerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveServersFromServerGroup(request)
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

// RemoveServersFromServerGroupWithCallback invokes the gwlb.RemoveServersFromServerGroup API asynchronously
func (client *Client) RemoveServersFromServerGroupWithCallback(request *RemoveServersFromServerGroupRequest, callback func(response *RemoveServersFromServerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveServersFromServerGroupResponse
		var err error
		defer close(result)
		response, err = client.RemoveServersFromServerGroup(request)
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

// RemoveServersFromServerGroupRequest is the request struct for api RemoveServersFromServerGroup
type RemoveServersFromServerGroupRequest struct {
	*requests.RpcRequest
	ClientToken   string                                 `position:"Body" name:"ClientToken"`
	ServerGroupId string                                 `position:"Body" name:"ServerGroupId"`
	Servers       *[]RemoveServersFromServerGroupServers `position:"Body" name:"Servers"  type:"Repeated"`
	DryRun        requests.Boolean                       `position:"Body" name:"DryRun"`
}

// RemoveServersFromServerGroupServers is a repeated param struct in RemoveServersFromServerGroupRequest
type RemoveServersFromServerGroupServers struct {
	Port       string `name:"Port"`
	ServerId   string `name:"ServerId"`
	ServerIp   string `name:"ServerIp"`
	ServerType string `name:"ServerType"`
}

// RemoveServersFromServerGroupResponse is the response struct for api RemoveServersFromServerGroup
type RemoveServersFromServerGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveServersFromServerGroupRequest creates a request to invoke RemoveServersFromServerGroup API
func CreateRemoveServersFromServerGroupRequest() (request *RemoveServersFromServerGroupRequest) {
	request = &RemoveServersFromServerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Gwlb", "2024-04-15", "RemoveServersFromServerGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateRemoveServersFromServerGroupResponse creates a response to parse from RemoveServersFromServerGroup response
func CreateRemoveServersFromServerGroupResponse() (response *RemoveServersFromServerGroupResponse) {
	response = &RemoveServersFromServerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

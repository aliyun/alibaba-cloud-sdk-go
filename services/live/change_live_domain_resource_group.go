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

// ChangeLiveDomainResourceGroup invokes the live.ChangeLiveDomainResourceGroup API synchronously
func (client *Client) ChangeLiveDomainResourceGroup(request *ChangeLiveDomainResourceGroupRequest) (response *ChangeLiveDomainResourceGroupResponse, err error) {
	response = CreateChangeLiveDomainResourceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ChangeLiveDomainResourceGroupWithChan invokes the live.ChangeLiveDomainResourceGroup API asynchronously
func (client *Client) ChangeLiveDomainResourceGroupWithChan(request *ChangeLiveDomainResourceGroupRequest) (<-chan *ChangeLiveDomainResourceGroupResponse, <-chan error) {
	responseChan := make(chan *ChangeLiveDomainResourceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangeLiveDomainResourceGroup(request)
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

// ChangeLiveDomainResourceGroupWithCallback invokes the live.ChangeLiveDomainResourceGroup API asynchronously
func (client *Client) ChangeLiveDomainResourceGroupWithCallback(request *ChangeLiveDomainResourceGroupRequest, callback func(response *ChangeLiveDomainResourceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangeLiveDomainResourceGroupResponse
		var err error
		defer close(result)
		response, err = client.ChangeLiveDomainResourceGroup(request)
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

// ChangeLiveDomainResourceGroupRequest is the request struct for api ChangeLiveDomainResourceGroup
type ChangeLiveDomainResourceGroupRequest struct {
	*requests.RpcRequest
	DomainName         string           `position:"Query" name:"DomainName"`
	OwnerId            requests.Integer `position:"Query" name:"OwnerId"`
	NewResourceGroupId string           `position:"Query" name:"NewResourceGroupId"`
}

// ChangeLiveDomainResourceGroupResponse is the response struct for api ChangeLiveDomainResourceGroup
type ChangeLiveDomainResourceGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateChangeLiveDomainResourceGroupRequest creates a request to invoke ChangeLiveDomainResourceGroup API
func CreateChangeLiveDomainResourceGroupRequest() (request *ChangeLiveDomainResourceGroupRequest) {
	request = &ChangeLiveDomainResourceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ChangeLiveDomainResourceGroup", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateChangeLiveDomainResourceGroupResponse creates a response to parse from ChangeLiveDomainResourceGroup response
func CreateChangeLiveDomainResourceGroupResponse() (response *ChangeLiveDomainResourceGroupResponse) {
	response = &ChangeLiveDomainResourceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

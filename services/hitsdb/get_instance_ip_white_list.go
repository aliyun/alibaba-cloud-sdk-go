package hitsdb

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

// GetInstanceIpWhiteList invokes the hitsdb.GetInstanceIpWhiteList API synchronously
func (client *Client) GetInstanceIpWhiteList(request *GetInstanceIpWhiteListRequest) (response *GetInstanceIpWhiteListResponse, err error) {
	response = CreateGetInstanceIpWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceIpWhiteListWithChan invokes the hitsdb.GetInstanceIpWhiteList API asynchronously
func (client *Client) GetInstanceIpWhiteListWithChan(request *GetInstanceIpWhiteListRequest) (<-chan *GetInstanceIpWhiteListResponse, <-chan error) {
	responseChan := make(chan *GetInstanceIpWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceIpWhiteList(request)
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

// GetInstanceIpWhiteListWithCallback invokes the hitsdb.GetInstanceIpWhiteList API asynchronously
func (client *Client) GetInstanceIpWhiteListWithCallback(request *GetInstanceIpWhiteListRequest, callback func(response *GetInstanceIpWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceIpWhiteListResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceIpWhiteList(request)
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

// GetInstanceIpWhiteListRequest is the request struct for api GetInstanceIpWhiteList
type GetInstanceIpWhiteListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	GroupName            string           `position:"Query" name:"GroupName"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// GetInstanceIpWhiteListResponse is the response struct for api GetInstanceIpWhiteList
type GetInstanceIpWhiteListResponse struct {
	*responses.BaseResponse
	InstanceId string          `json:"InstanceId" xml:"InstanceId"`
	RequestId  string          `json:"RequestId" xml:"RequestId"`
	IpList     []string        `json:"IpList" xml:"IpList"`
	GroupList  []GroupListItem `json:"GroupList" xml:"GroupList"`
}

// CreateGetInstanceIpWhiteListRequest creates a request to invoke GetInstanceIpWhiteList API
func CreateGetInstanceIpWhiteListRequest() (request *GetInstanceIpWhiteListRequest) {
	request = &GetInstanceIpWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hitsdb", "2020-06-15", "GetInstanceIpWhiteList", "hitsdb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetInstanceIpWhiteListResponse creates a response to parse from GetInstanceIpWhiteList response
func CreateGetInstanceIpWhiteListResponse() (response *GetInstanceIpWhiteListResponse) {
	response = &GetInstanceIpWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package polardb

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

// ModifyGlobalDatabaseNetwork invokes the polardb.ModifyGlobalDatabaseNetwork API synchronously
func (client *Client) ModifyGlobalDatabaseNetwork(request *ModifyGlobalDatabaseNetworkRequest) (response *ModifyGlobalDatabaseNetworkResponse, err error) {
	response = CreateModifyGlobalDatabaseNetworkResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyGlobalDatabaseNetworkWithChan invokes the polardb.ModifyGlobalDatabaseNetwork API asynchronously
func (client *Client) ModifyGlobalDatabaseNetworkWithChan(request *ModifyGlobalDatabaseNetworkRequest) (<-chan *ModifyGlobalDatabaseNetworkResponse, <-chan error) {
	responseChan := make(chan *ModifyGlobalDatabaseNetworkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyGlobalDatabaseNetwork(request)
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

// ModifyGlobalDatabaseNetworkWithCallback invokes the polardb.ModifyGlobalDatabaseNetwork API asynchronously
func (client *Client) ModifyGlobalDatabaseNetworkWithCallback(request *ModifyGlobalDatabaseNetworkRequest, callback func(response *ModifyGlobalDatabaseNetworkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyGlobalDatabaseNetworkResponse
		var err error
		defer close(result)
		response, err = client.ModifyGlobalDatabaseNetwork(request)
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

// ModifyGlobalDatabaseNetworkRequest is the request struct for api ModifyGlobalDatabaseNetwork
type ModifyGlobalDatabaseNetworkRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId        string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken          string           `position:"Query" name:"SecurityToken"`
	GDNId                  string           `position:"Query" name:"GDNId"`
	GDNDescription         string           `position:"Query" name:"GDNDescription"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	EnableGlobalDomainName requests.Boolean `position:"Query" name:"EnableGlobalDomainName"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyGlobalDatabaseNetworkResponse is the response struct for api ModifyGlobalDatabaseNetwork
type ModifyGlobalDatabaseNetworkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyGlobalDatabaseNetworkRequest creates a request to invoke ModifyGlobalDatabaseNetwork API
func CreateModifyGlobalDatabaseNetworkRequest() (request *ModifyGlobalDatabaseNetworkRequest) {
	request = &ModifyGlobalDatabaseNetworkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyGlobalDatabaseNetwork", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyGlobalDatabaseNetworkResponse creates a response to parse from ModifyGlobalDatabaseNetwork response
func CreateModifyGlobalDatabaseNetworkResponse() (response *ModifyGlobalDatabaseNetworkResponse) {
	response = &ModifyGlobalDatabaseNetworkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

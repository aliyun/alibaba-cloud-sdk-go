package rds

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

// ModifyDBProxyEndpointAddress invokes the rds.ModifyDBProxyEndpointAddress API synchronously
func (client *Client) ModifyDBProxyEndpointAddress(request *ModifyDBProxyEndpointAddressRequest) (response *ModifyDBProxyEndpointAddressResponse, err error) {
	response = CreateModifyDBProxyEndpointAddressResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBProxyEndpointAddressWithChan invokes the rds.ModifyDBProxyEndpointAddress API asynchronously
func (client *Client) ModifyDBProxyEndpointAddressWithChan(request *ModifyDBProxyEndpointAddressRequest) (<-chan *ModifyDBProxyEndpointAddressResponse, <-chan error) {
	responseChan := make(chan *ModifyDBProxyEndpointAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBProxyEndpointAddress(request)
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

// ModifyDBProxyEndpointAddressWithCallback invokes the rds.ModifyDBProxyEndpointAddress API asynchronously
func (client *Client) ModifyDBProxyEndpointAddressWithCallback(request *ModifyDBProxyEndpointAddressRequest, callback func(response *ModifyDBProxyEndpointAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBProxyEndpointAddressResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBProxyEndpointAddress(request)
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

// ModifyDBProxyEndpointAddressRequest is the request struct for api ModifyDBProxyEndpointAddress
type ModifyDBProxyEndpointAddressRequest struct {
	*requests.RpcRequest
	ResourceOwnerId             requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBProxyConnectStringNetType string           `position:"Query" name:"DBProxyConnectStringNetType"`
	DBInstanceId                string           `position:"Query" name:"DBInstanceId"`
	DBProxyNewConnectStringPort string           `position:"Query" name:"DBProxyNewConnectStringPort"`
	ResourceOwnerAccount        string           `position:"Query" name:"ResourceOwnerAccount"`
	DBProxyEngineType           string           `position:"Query" name:"DBProxyEngineType"`
	OwnerId                     requests.Integer `position:"Query" name:"OwnerId"`
	DBProxyEndpointId           string           `position:"Query" name:"DBProxyEndpointId"`
	DBProxyNewConnectString     string           `position:"Query" name:"DBProxyNewConnectString"`
}

// ModifyDBProxyEndpointAddressResponse is the response struct for api ModifyDBProxyEndpointAddress
type ModifyDBProxyEndpointAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBProxyEndpointAddressRequest creates a request to invoke ModifyDBProxyEndpointAddress API
func CreateModifyDBProxyEndpointAddressRequest() (request *ModifyDBProxyEndpointAddressRequest) {
	request = &ModifyDBProxyEndpointAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBProxyEndpointAddress", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyDBProxyEndpointAddressResponse creates a response to parse from ModifyDBProxyEndpointAddress response
func CreateModifyDBProxyEndpointAddressResponse() (response *ModifyDBProxyEndpointAddressResponse) {
	response = &ModifyDBProxyEndpointAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

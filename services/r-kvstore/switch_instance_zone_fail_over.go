package r_kvstore

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

// SwitchInstanceZoneFailOver invokes the r_kvstore.SwitchInstanceZoneFailOver API synchronously
func (client *Client) SwitchInstanceZoneFailOver(request *SwitchInstanceZoneFailOverRequest) (response *SwitchInstanceZoneFailOverResponse, err error) {
	response = CreateSwitchInstanceZoneFailOverResponse()
	err = client.DoAction(request, response)
	return
}

// SwitchInstanceZoneFailOverWithChan invokes the r_kvstore.SwitchInstanceZoneFailOver API asynchronously
func (client *Client) SwitchInstanceZoneFailOverWithChan(request *SwitchInstanceZoneFailOverRequest) (<-chan *SwitchInstanceZoneFailOverResponse, <-chan error) {
	responseChan := make(chan *SwitchInstanceZoneFailOverResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SwitchInstanceZoneFailOver(request)
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

// SwitchInstanceZoneFailOverWithCallback invokes the r_kvstore.SwitchInstanceZoneFailOver API asynchronously
func (client *Client) SwitchInstanceZoneFailOverWithCallback(request *SwitchInstanceZoneFailOverRequest, callback func(response *SwitchInstanceZoneFailOverResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SwitchInstanceZoneFailOverResponse
		var err error
		defer close(result)
		response, err = client.SwitchInstanceZoneFailOver(request)
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

// SwitchInstanceZoneFailOverRequest is the request struct for api SwitchInstanceZoneFailOver
type SwitchInstanceZoneFailOverRequest struct {
	*requests.RpcRequest
	SiteFaultTime string `position:"Query" name:"SiteFaultTime"`
	TargetZoneId  string `position:"Query" name:"TargetZoneId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
}

// SwitchInstanceZoneFailOverResponse is the response struct for api SwitchInstanceZoneFailOver
type SwitchInstanceZoneFailOverResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSwitchInstanceZoneFailOverRequest creates a request to invoke SwitchInstanceZoneFailOver API
func CreateSwitchInstanceZoneFailOverRequest() (request *SwitchInstanceZoneFailOverRequest) {
	request = &SwitchInstanceZoneFailOverRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "SwitchInstanceZoneFailOver", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSwitchInstanceZoneFailOverResponse creates a response to parse from SwitchInstanceZoneFailOver response
func CreateSwitchInstanceZoneFailOverResponse() (response *SwitchInstanceZoneFailOverResponse) {
	response = &SwitchInstanceZoneFailOverResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

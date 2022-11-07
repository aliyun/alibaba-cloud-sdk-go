package cbn

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

// ListGrantVSwitchesToCen invokes the cbn.ListGrantVSwitchesToCen API synchronously
func (client *Client) ListGrantVSwitchesToCen(request *ListGrantVSwitchesToCenRequest) (response *ListGrantVSwitchesToCenResponse, err error) {
	response = CreateListGrantVSwitchesToCenResponse()
	err = client.DoAction(request, response)
	return
}

// ListGrantVSwitchesToCenWithChan invokes the cbn.ListGrantVSwitchesToCen API asynchronously
func (client *Client) ListGrantVSwitchesToCenWithChan(request *ListGrantVSwitchesToCenRequest) (<-chan *ListGrantVSwitchesToCenResponse, <-chan error) {
	responseChan := make(chan *ListGrantVSwitchesToCenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListGrantVSwitchesToCen(request)
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

// ListGrantVSwitchesToCenWithCallback invokes the cbn.ListGrantVSwitchesToCen API asynchronously
func (client *Client) ListGrantVSwitchesToCenWithCallback(request *ListGrantVSwitchesToCenRequest, callback func(response *ListGrantVSwitchesToCenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListGrantVSwitchesToCenResponse
		var err error
		defer close(result)
		response, err = client.ListGrantVSwitchesToCen(request)
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

// ListGrantVSwitchesToCenRequest is the request struct for api ListGrantVSwitchesToCen
type ListGrantVSwitchesToCenRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                string           `position:"Query" name:"CenId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VpcId                string           `position:"Query" name:"VpcId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// ListGrantVSwitchesToCenResponse is the response struct for api ListGrantVSwitchesToCen
type ListGrantVSwitchesToCenResponse struct {
	*responses.BaseResponse
	PageSize   int       `json:"PageSize" xml:"PageSize"`
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	PageNumber int       `json:"PageNumber" xml:"PageNumber"`
	TotalCount int       `json:"TotalCount" xml:"TotalCount"`
	VSwitches  []VSwitch `json:"VSwitches" xml:"VSwitches"`
}

// CreateListGrantVSwitchesToCenRequest creates a request to invoke ListGrantVSwitchesToCen API
func CreateListGrantVSwitchesToCenRequest() (request *ListGrantVSwitchesToCenRequest) {
	request = &ListGrantVSwitchesToCenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "ListGrantVSwitchesToCen", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListGrantVSwitchesToCenResponse creates a response to parse from ListGrantVSwitchesToCen response
func CreateListGrantVSwitchesToCenResponse() (response *ListGrantVSwitchesToCenResponse) {
	response = &ListGrantVSwitchesToCenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

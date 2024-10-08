package biz_ipdb

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

// AddLocationAttribute invokes the biz_ipdb.AddLocationAttribute API synchronously
func (client *Client) AddLocationAttribute(request *AddLocationAttributeRequest) (response *AddLocationAttributeResponse, err error) {
	response = CreateAddLocationAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// AddLocationAttributeWithChan invokes the biz_ipdb.AddLocationAttribute API asynchronously
func (client *Client) AddLocationAttributeWithChan(request *AddLocationAttributeRequest) (<-chan *AddLocationAttributeResponse, <-chan error) {
	responseChan := make(chan *AddLocationAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLocationAttribute(request)
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

// AddLocationAttributeWithCallback invokes the biz_ipdb.AddLocationAttribute API asynchronously
func (client *Client) AddLocationAttributeWithCallback(request *AddLocationAttributeRequest, callback func(response *AddLocationAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLocationAttributeResponse
		var err error
		defer close(result)
		response, err = client.AddLocationAttribute(request)
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

// AddLocationAttributeRequest is the request struct for api AddLocationAttribute
type AddLocationAttributeRequest struct {
	*requests.RpcRequest
	LocationNo   string `position:"Query" name:"LocationNo"`
	LocationType string `position:"Query" name:"LocationType"`
	Value        string `position:"Query" name:"Value"`
	BizLine      string `position:"Query" name:"BizLine"`
	Key          string `position:"Query" name:"Key"`
}

// AddLocationAttributeResponse is the response struct for api AddLocationAttribute
type AddLocationAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddLocationAttributeRequest creates a request to invoke AddLocationAttribute API
func CreateAddLocationAttributeRequest() (request *AddLocationAttributeRequest) {
	request = &AddLocationAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Biz-ipdb", "2016-08-08", "AddLocationAttribute", "", "")
	request.Method = requests.POST
	return
}

// CreateAddLocationAttributeResponse creates a response to parse from AddLocationAttribute response
func CreateAddLocationAttributeResponse() (response *AddLocationAttributeResponse) {
	response = &AddLocationAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

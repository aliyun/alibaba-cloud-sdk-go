package scdn

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

// SetScdnCcInfo invokes the scdn.SetScdnCcInfo API synchronously
func (client *Client) SetScdnCcInfo(request *SetScdnCcInfoRequest) (response *SetScdnCcInfoResponse, err error) {
	response = CreateSetScdnCcInfoResponse()
	err = client.DoAction(request, response)
	return
}

// SetScdnCcInfoWithChan invokes the scdn.SetScdnCcInfo API asynchronously
func (client *Client) SetScdnCcInfoWithChan(request *SetScdnCcInfoRequest) (<-chan *SetScdnCcInfoResponse, <-chan error) {
	responseChan := make(chan *SetScdnCcInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetScdnCcInfo(request)
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

// SetScdnCcInfoWithCallback invokes the scdn.SetScdnCcInfo API asynchronously
func (client *Client) SetScdnCcInfoWithCallback(request *SetScdnCcInfoRequest, callback func(response *SetScdnCcInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetScdnCcInfoResponse
		var err error
		defer close(result)
		response, err = client.SetScdnCcInfo(request)
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

// SetScdnCcInfoRequest is the request struct for api SetScdnCcInfo
type SetScdnCcInfoRequest struct {
	*requests.RpcRequest
	Status string `position:"Query" name:"Status"`
}

// SetScdnCcInfoResponse is the response struct for api SetScdnCcInfo
type SetScdnCcInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetScdnCcInfoRequest creates a request to invoke SetScdnCcInfo API
func CreateSetScdnCcInfoRequest() (request *SetScdnCcInfoRequest) {
	request = &SetScdnCcInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "SetScdnCcInfo", "", "")
	request.Method = requests.GET
	return
}

// CreateSetScdnCcInfoResponse creates a response to parse from SetScdnCcInfo response
func CreateSetScdnCcInfoResponse() (response *SetScdnCcInfoResponse) {
	response = &SetScdnCcInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

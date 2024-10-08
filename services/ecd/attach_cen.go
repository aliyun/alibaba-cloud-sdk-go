package ecd

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

// AttachCen invokes the ecd.AttachCen API synchronously
func (client *Client) AttachCen(request *AttachCenRequest) (response *AttachCenResponse, err error) {
	response = CreateAttachCenResponse()
	err = client.DoAction(request, response)
	return
}

// AttachCenWithChan invokes the ecd.AttachCen API asynchronously
func (client *Client) AttachCenWithChan(request *AttachCenRequest) (<-chan *AttachCenResponse, <-chan error) {
	responseChan := make(chan *AttachCenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachCen(request)
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

// AttachCenWithCallback invokes the ecd.AttachCen API asynchronously
func (client *Client) AttachCenWithCallback(request *AttachCenRequest, callback func(response *AttachCenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachCenResponse
		var err error
		defer close(result)
		response, err = client.AttachCen(request)
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

// AttachCenRequest is the request struct for api AttachCen
type AttachCenRequest struct {
	*requests.RpcRequest
	OfficeSiteId string           `position:"Query" name:"OfficeSiteId"`
	CenId        string           `position:"Query" name:"CenId"`
	CenOwnerId   requests.Integer `position:"Query" name:"CenOwnerId"`
	VerifyCode   string           `position:"Query" name:"VerifyCode"`
}

// AttachCenResponse is the response struct for api AttachCen
type AttachCenResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachCenRequest creates a request to invoke AttachCen API
func CreateAttachCenRequest() (request *AttachCenRequest) {
	request = &AttachCenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "AttachCen", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachCenResponse creates a response to parse from AttachCen response
func CreateAttachCenResponse() (response *AttachCenResponse) {
	response = &AttachCenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

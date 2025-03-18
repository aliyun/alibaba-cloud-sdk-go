package ens

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

// ModifySnatEntry invokes the ens.ModifySnatEntry API synchronously
func (client *Client) ModifySnatEntry(request *ModifySnatEntryRequest) (response *ModifySnatEntryResponse, err error) {
	response = CreateModifySnatEntryResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySnatEntryWithChan invokes the ens.ModifySnatEntry API asynchronously
func (client *Client) ModifySnatEntryWithChan(request *ModifySnatEntryRequest) (<-chan *ModifySnatEntryResponse, <-chan error) {
	responseChan := make(chan *ModifySnatEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySnatEntry(request)
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

// ModifySnatEntryWithCallback invokes the ens.ModifySnatEntry API asynchronously
func (client *Client) ModifySnatEntryWithCallback(request *ModifySnatEntryRequest, callback func(response *ModifySnatEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySnatEntryResponse
		var err error
		defer close(result)
		response, err = client.ModifySnatEntry(request)
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

// ModifySnatEntryRequest is the request struct for api ModifySnatEntry
type ModifySnatEntryRequest struct {
	*requests.RpcRequest
	SnatIp        string           `position:"Query" name:"SnatIp"`
	EipAffinity   requests.Boolean `position:"Query" name:"EipAffinity"`
	SnatEntryId   string           `position:"Query" name:"SnatEntryId"`
	SnatEntryName string           `position:"Query" name:"SnatEntryName"`
	IspAffinity   requests.Boolean `position:"Query" name:"IspAffinity"`
}

// ModifySnatEntryResponse is the response struct for api ModifySnatEntry
type ModifySnatEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySnatEntryRequest creates a request to invoke ModifySnatEntry API
func CreateModifySnatEntryRequest() (request *ModifySnatEntryRequest) {
	request = &ModifySnatEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "ModifySnatEntry", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySnatEntryResponse creates a response to parse from ModifySnatEntry response
func CreateModifySnatEntryResponse() (response *ModifySnatEntryResponse) {
	response = &ModifySnatEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package drds

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

// RemoveDrdsDb invokes the drds.RemoveDrdsDb API synchronously
func (client *Client) RemoveDrdsDb(request *RemoveDrdsDbRequest) (response *RemoveDrdsDbResponse, err error) {
	response = CreateRemoveDrdsDbResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveDrdsDbWithChan invokes the drds.RemoveDrdsDb API asynchronously
func (client *Client) RemoveDrdsDbWithChan(request *RemoveDrdsDbRequest) (<-chan *RemoveDrdsDbResponse, <-chan error) {
	responseChan := make(chan *RemoveDrdsDbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveDrdsDb(request)
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

// RemoveDrdsDbWithCallback invokes the drds.RemoveDrdsDb API asynchronously
func (client *Client) RemoveDrdsDbWithCallback(request *RemoveDrdsDbRequest, callback func(response *RemoveDrdsDbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveDrdsDbResponse
		var err error
		defer close(result)
		response, err = client.RemoveDrdsDb(request)
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

// RemoveDrdsDbRequest is the request struct for api RemoveDrdsDb
type RemoveDrdsDbRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
}

// RemoveDrdsDbResponse is the response struct for api RemoveDrdsDb
type RemoveDrdsDbResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveDrdsDbRequest creates a request to invoke RemoveDrdsDb API
func CreateRemoveDrdsDbRequest() (request *RemoveDrdsDbRequest) {
	request = &RemoveDrdsDbRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "RemoveDrdsDb", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveDrdsDbResponse creates a response to parse from RemoveDrdsDb response
func CreateRemoveDrdsDbResponse() (response *RemoveDrdsDbResponse) {
	response = &RemoveDrdsDbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package pairecservice

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

// DeleteCrowd invokes the pairecservice.DeleteCrowd API synchronously
func (client *Client) DeleteCrowd(request *DeleteCrowdRequest) (response *DeleteCrowdResponse, err error) {
	response = CreateDeleteCrowdResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCrowdWithChan invokes the pairecservice.DeleteCrowd API asynchronously
func (client *Client) DeleteCrowdWithChan(request *DeleteCrowdRequest) (<-chan *DeleteCrowdResponse, <-chan error) {
	responseChan := make(chan *DeleteCrowdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCrowd(request)
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

// DeleteCrowdWithCallback invokes the pairecservice.DeleteCrowd API asynchronously
func (client *Client) DeleteCrowdWithCallback(request *DeleteCrowdRequest, callback func(response *DeleteCrowdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCrowdResponse
		var err error
		defer close(result)
		response, err = client.DeleteCrowd(request)
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

// DeleteCrowdRequest is the request struct for api DeleteCrowd
type DeleteCrowdRequest struct {
	*requests.RoaRequest
	CrowdId    string `position:"Path" name:"CrowdId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DeleteCrowdResponse is the response struct for api DeleteCrowd
type DeleteCrowdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCrowdRequest creates a request to invoke DeleteCrowd API
func CreateDeleteCrowdRequest() (request *DeleteCrowdRequest) {
	request = &DeleteCrowdRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "DeleteCrowd", "/api/v1/crowds/[CrowdId]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteCrowdResponse creates a response to parse from DeleteCrowd response
func CreateDeleteCrowdResponse() (response *DeleteCrowdResponse) {
	response = &DeleteCrowdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

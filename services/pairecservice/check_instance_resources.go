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

// CheckInstanceResources invokes the pairecservice.CheckInstanceResources API synchronously
func (client *Client) CheckInstanceResources(request *CheckInstanceResourcesRequest) (response *CheckInstanceResourcesResponse, err error) {
	response = CreateCheckInstanceResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// CheckInstanceResourcesWithChan invokes the pairecservice.CheckInstanceResources API asynchronously
func (client *Client) CheckInstanceResourcesWithChan(request *CheckInstanceResourcesRequest) (<-chan *CheckInstanceResourcesResponse, <-chan error) {
	responseChan := make(chan *CheckInstanceResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckInstanceResources(request)
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

// CheckInstanceResourcesWithCallback invokes the pairecservice.CheckInstanceResources API asynchronously
func (client *Client) CheckInstanceResourcesWithCallback(request *CheckInstanceResourcesRequest, callback func(response *CheckInstanceResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckInstanceResourcesResponse
		var err error
		defer close(result)
		response, err = client.CheckInstanceResources(request)
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

// CheckInstanceResourcesRequest is the request struct for api CheckInstanceResources
type CheckInstanceResourcesRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
	Body       string `position:"Body" name:"body"`
}

// CheckInstanceResourcesResponse is the response struct for api CheckInstanceResources
type CheckInstanceResourcesResponse struct {
	*responses.BaseResponse
	RequestId string          `json:"RequestId" xml:"RequestId"`
	Resources []ResourcesItem `json:"Resources" xml:"Resources"`
}

// CreateCheckInstanceResourcesRequest creates a request to invoke CheckInstanceResources API
func CreateCheckInstanceResourcesRequest() (request *CheckInstanceResourcesRequest) {
	request = &CheckInstanceResourcesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "CheckInstanceResources", "/api/v1/instances/[InstanceId]/action/checkresources", "", "")
	request.Method = requests.POST
	return
}

// CreateCheckInstanceResourcesResponse creates a response to parse from CheckInstanceResources response
func CreateCheckInstanceResourcesResponse() (response *CheckInstanceResourcesResponse) {
	response = &CheckInstanceResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

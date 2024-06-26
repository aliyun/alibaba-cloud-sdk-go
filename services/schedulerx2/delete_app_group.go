package schedulerx2

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

// DeleteAppGroup invokes the schedulerx2.DeleteAppGroup API synchronously
func (client *Client) DeleteAppGroup(request *DeleteAppGroupRequest) (response *DeleteAppGroupResponse, err error) {
	response = CreateDeleteAppGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAppGroupWithChan invokes the schedulerx2.DeleteAppGroup API asynchronously
func (client *Client) DeleteAppGroupWithChan(request *DeleteAppGroupRequest) (<-chan *DeleteAppGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteAppGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAppGroup(request)
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

// DeleteAppGroupWithCallback invokes the schedulerx2.DeleteAppGroup API asynchronously
func (client *Client) DeleteAppGroupWithCallback(request *DeleteAppGroupRequest, callback func(response *DeleteAppGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAppGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteAppGroup(request)
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

// DeleteAppGroupRequest is the request struct for api DeleteAppGroup
type DeleteAppGroupRequest struct {
	*requests.RpcRequest
	GroupId    string           `position:"Query" name:"GroupId"`
	DeleteJobs requests.Boolean `position:"Query" name:"DeleteJobs"`
	Namespace  string           `position:"Query" name:"Namespace"`
}

// DeleteAppGroupResponse is the response struct for api DeleteAppGroup
type DeleteAppGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteAppGroupRequest creates a request to invoke DeleteAppGroup API
func CreateDeleteAppGroupRequest() (request *DeleteAppGroupRequest) {
	request = &DeleteAppGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "DeleteAppGroup", "schedulerx2", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteAppGroupResponse creates a response to parse from DeleteAppGroup response
func CreateDeleteAppGroupResponse() (response *DeleteAppGroupResponse) {
	response = &DeleteAppGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

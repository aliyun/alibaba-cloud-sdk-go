package ehpc

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

// DeleteUsers invokes the ehpc.DeleteUsers API synchronously
func (client *Client) DeleteUsers(request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
	response = CreateDeleteUsersResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteUsersWithChan invokes the ehpc.DeleteUsers API asynchronously
func (client *Client) DeleteUsersWithChan(request *DeleteUsersRequest) (<-chan *DeleteUsersResponse, <-chan error) {
	responseChan := make(chan *DeleteUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteUsers(request)
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

// DeleteUsersWithCallback invokes the ehpc.DeleteUsers API asynchronously
func (client *Client) DeleteUsersWithCallback(request *DeleteUsersRequest, callback func(response *DeleteUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteUsersResponse
		var err error
		defer close(result)
		response, err = client.DeleteUsers(request)
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

// DeleteUsersRequest is the request struct for api DeleteUsers
type DeleteUsersRequest struct {
	*requests.RpcRequest
	ClusterId string             `position:"Query" name:"ClusterId"`
	Async     requests.Boolean   `position:"Query" name:"Async"`
	User      *[]DeleteUsersUser `position:"Query" name:"User"  type:"Repeated"`
}

// DeleteUsersUser is a repeated param struct in DeleteUsersRequest
type DeleteUsersUser struct {
	Name string `name:"Name"`
}

// DeleteUsersResponse is the response struct for api DeleteUsers
type DeleteUsersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteUsersRequest creates a request to invoke DeleteUsers API
func CreateDeleteUsersRequest() (request *DeleteUsersRequest) {
	request = &DeleteUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "DeleteUsers", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDeleteUsersResponse creates a response to parse from DeleteUsers response
func CreateDeleteUsersResponse() (response *DeleteUsersResponse) {
	response = &DeleteUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package domain

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

// QueryServerLock invokes the domain.QueryServerLock API synchronously
func (client *Client) QueryServerLock(request *QueryServerLockRequest) (response *QueryServerLockResponse, err error) {
	response = CreateQueryServerLockResponse()
	err = client.DoAction(request, response)
	return
}

// QueryServerLockWithChan invokes the domain.QueryServerLock API asynchronously
func (client *Client) QueryServerLockWithChan(request *QueryServerLockRequest) (<-chan *QueryServerLockResponse, <-chan error) {
	responseChan := make(chan *QueryServerLockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryServerLock(request)
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

// QueryServerLockWithCallback invokes the domain.QueryServerLock API asynchronously
func (client *Client) QueryServerLockWithCallback(request *QueryServerLockRequest, callback func(response *QueryServerLockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryServerLockResponse
		var err error
		defer close(result)
		response, err = client.QueryServerLock(request)
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

// QueryServerLockRequest is the request struct for api QueryServerLock
type QueryServerLockRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// QueryServerLockResponse is the response struct for api QueryServerLock
type QueryServerLockResponse struct {
	*responses.BaseResponse
	StartDate        string `json:"StartDate" xml:"StartDate"`
	GmtCreate        string `json:"GmtCreate" xml:"GmtCreate"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	ExpireDate       string `json:"ExpireDate" xml:"ExpireDate"`
	DomainName       string `json:"DomainName" xml:"DomainName"`
	UserId           string `json:"UserId" xml:"UserId"`
	GmtModified      string `json:"GmtModified" xml:"GmtModified"`
	DomainInstanceId string `json:"DomainInstanceId" xml:"DomainInstanceId"`
	LockInstanceId   string `json:"LockInstanceId" xml:"LockInstanceId"`
	ServerLockStatus int    `json:"ServerLockStatus" xml:"ServerLockStatus"`
	LockProductId    string `json:"LockProductId" xml:"LockProductId"`
}

// CreateQueryServerLockRequest creates a request to invoke QueryServerLock API
func CreateQueryServerLockRequest() (request *QueryServerLockRequest) {
	request = &QueryServerLockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryServerLock", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryServerLockResponse creates a response to parse from QueryServerLock response
func CreateQueryServerLockResponse() (response *QueryServerLockResponse) {
	response = &QueryServerLockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

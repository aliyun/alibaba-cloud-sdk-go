package polardb

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

// CheckAccountName invokes the polardb.CheckAccountName API synchronously
func (client *Client) CheckAccountName(request *CheckAccountNameRequest) (response *CheckAccountNameResponse, err error) {
	response = CreateCheckAccountNameResponse()
	err = client.DoAction(request, response)
	return
}

// CheckAccountNameWithChan invokes the polardb.CheckAccountName API asynchronously
func (client *Client) CheckAccountNameWithChan(request *CheckAccountNameRequest) (<-chan *CheckAccountNameResponse, <-chan error) {
	responseChan := make(chan *CheckAccountNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckAccountName(request)
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

// CheckAccountNameWithCallback invokes the polardb.CheckAccountName API asynchronously
func (client *Client) CheckAccountNameWithCallback(request *CheckAccountNameRequest, callback func(response *CheckAccountNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckAccountNameResponse
		var err error
		defer close(result)
		response, err = client.CheckAccountName(request)
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

// CheckAccountNameRequest is the request struct for api CheckAccountName
type CheckAccountNameRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountName          string           `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CheckAccountNameResponse is the response struct for api CheckAccountName
type CheckAccountNameResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCheckAccountNameRequest creates a request to invoke CheckAccountName API
func CreateCheckAccountNameRequest() (request *CheckAccountNameRequest) {
	request = &CheckAccountNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "CheckAccountName", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckAccountNameResponse creates a response to parse from CheckAccountName response
func CreateCheckAccountNameResponse() (response *CheckAccountNameResponse) {
	response = &CheckAccountNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

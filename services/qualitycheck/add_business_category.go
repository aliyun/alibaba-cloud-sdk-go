package qualitycheck

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

// AddBusinessCategory invokes the qualitycheck.AddBusinessCategory API synchronously
func (client *Client) AddBusinessCategory(request *AddBusinessCategoryRequest) (response *AddBusinessCategoryResponse, err error) {
	response = CreateAddBusinessCategoryResponse()
	err = client.DoAction(request, response)
	return
}

// AddBusinessCategoryWithChan invokes the qualitycheck.AddBusinessCategory API asynchronously
func (client *Client) AddBusinessCategoryWithChan(request *AddBusinessCategoryRequest) (<-chan *AddBusinessCategoryResponse, <-chan error) {
	responseChan := make(chan *AddBusinessCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddBusinessCategory(request)
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

// AddBusinessCategoryWithCallback invokes the qualitycheck.AddBusinessCategory API asynchronously
func (client *Client) AddBusinessCategoryWithCallback(request *AddBusinessCategoryRequest, callback func(response *AddBusinessCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddBusinessCategoryResponse
		var err error
		defer close(result)
		response, err = client.AddBusinessCategory(request)
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

// AddBusinessCategoryRequest is the request struct for api AddBusinessCategory
type AddBusinessCategoryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
	BaseMeAgentId   requests.Integer `position:"Query" name:"BaseMeAgentId"`
}

// AddBusinessCategoryResponse is the response struct for api AddBusinessCategory
type AddBusinessCategoryResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateAddBusinessCategoryRequest creates a request to invoke AddBusinessCategory API
func CreateAddBusinessCategoryRequest() (request *AddBusinessCategoryRequest) {
	request = &AddBusinessCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "AddBusinessCategory", "", "")
	request.Method = requests.POST
	return
}

// CreateAddBusinessCategoryResponse creates a response to parse from AddBusinessCategory response
func CreateAddBusinessCategoryResponse() (response *AddBusinessCategoryResponse) {
	response = &AddBusinessCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

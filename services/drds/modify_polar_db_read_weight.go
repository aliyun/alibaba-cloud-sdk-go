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

// ModifyPolarDbReadWeight invokes the drds.ModifyPolarDbReadWeight API synchronously
func (client *Client) ModifyPolarDbReadWeight(request *ModifyPolarDbReadWeightRequest) (response *ModifyPolarDbReadWeightResponse, err error) {
	response = CreateModifyPolarDbReadWeightResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPolarDbReadWeightWithChan invokes the drds.ModifyPolarDbReadWeight API asynchronously
func (client *Client) ModifyPolarDbReadWeightWithChan(request *ModifyPolarDbReadWeightRequest) (<-chan *ModifyPolarDbReadWeightResponse, <-chan error) {
	responseChan := make(chan *ModifyPolarDbReadWeightResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPolarDbReadWeight(request)
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

// ModifyPolarDbReadWeightWithCallback invokes the drds.ModifyPolarDbReadWeight API asynchronously
func (client *Client) ModifyPolarDbReadWeightWithCallback(request *ModifyPolarDbReadWeightRequest, callback func(response *ModifyPolarDbReadWeightResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPolarDbReadWeightResponse
		var err error
		defer close(result)
		response, err = client.ModifyPolarDbReadWeight(request)
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

// ModifyPolarDbReadWeightRequest is the request struct for api ModifyPolarDbReadWeight
type ModifyPolarDbReadWeightRequest struct {
	*requests.RpcRequest
	Weights        string `position:"Query" name:"Weights"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbNodeIds      string `position:"Query" name:"DbNodeIds"`
	DbName         string `position:"Query" name:"DbName"`
	DbInstanceId   string `position:"Query" name:"DbInstanceId"`
}

// ModifyPolarDbReadWeightResponse is the response struct for api ModifyPolarDbReadWeight
type ModifyPolarDbReadWeightResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyPolarDbReadWeightRequest creates a request to invoke ModifyPolarDbReadWeight API
func CreateModifyPolarDbReadWeightRequest() (request *ModifyPolarDbReadWeightRequest) {
	request = &ModifyPolarDbReadWeightRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "ModifyPolarDbReadWeight", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyPolarDbReadWeightResponse creates a response to parse from ModifyPolarDbReadWeight response
func CreateModifyPolarDbReadWeightResponse() (response *ModifyPolarDbReadWeightResponse) {
	response = &ModifyPolarDbReadWeightResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
